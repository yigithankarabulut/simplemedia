package apiserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	commentstorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/comment"
	friendsstorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/friends"
	likesstorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/likes"
	poststorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/post"
	userstorage "github.com/yigithankarabulut/simplemedia/src/internal/repository/user"
	commentservice "github.com/yigithankarabulut/simplemedia/src/internal/service/comment"
	friendsservice "github.com/yigithankarabulut/simplemedia/src/internal/service/friends"
	likesservice "github.com/yigithankarabulut/simplemedia/src/internal/service/likes"
	postservice "github.com/yigithankarabulut/simplemedia/src/internal/service/post"
	usersevice "github.com/yigithankarabulut/simplemedia/src/internal/service/user"
	commenthttphandler "github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/comment"
	friendshttphandler "github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/friends"
	likeshttphandler "github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/likes"
	posthttphandler "github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/post"
	"github.com/yigithankarabulut/simplemedia/src/internal/transport/httphandler/user"
	"github.com/yigithankarabulut/simplemedia/src/pkg/postgresql"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func New(opts ...Option) error {
	apiserver := &apiServer{
		logLevel: slog.LevelInfo,
		app:      fiber.New(),
		port:     ":8080",
	}
	for _, opt := range opts {
		opt(apiserver)
	}

	if apiserver.logger == nil {
		logHandlerOpts := &slog.HandlerOptions{Level: apiserver.logLevel}
		logHandler := slog.NewJSONHandler(os.Stdout, logHandlerOpts)
		apiserver.logger = slog.New(logHandler)
	}
	slog.SetDefault(apiserver.logger)
	logger := apiserver.logger

	db, err := postgresql.Connect()
	if err != nil {
		return err
	}

	AppendLayers(apiserver, db)

	shutdown := make(chan os.Signal, 1)
	apiError := make(chan error, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("starting api server...")
		apiError <- apiserver.app.Listen(apiserver.port)
	}()

	select {
	case err := <-apiError:
		return fmt.Errorf("listening error: %w", err)
	case <-shutdown:
		logger.Info("starting shutdown", "pid", os.Getpid())
		time.Sleep(1 * time.Second)
		defer logger.Info("shutdown completed", "pid", os.Getpid())
	}
	return nil
}

func AppendLayers(apiserver *apiServer, db *gorm.DB) {
	userRepository := userstorage.NewUserRepository(userstorage.WithUserRepositoryDB(db))
	postsRepository := poststorage.NewPostRepository(poststorage.WithPostRepositoryDB(db))
	commentRepository := commentstorage.NewCommentRepository(commentstorage.WithCommentRepositoryDB(db))
	likesRepository := likesstorage.NewLikeRepository(likesstorage.WithLikeRepositoryDB(db))
	friendsRepository := friendsstorage.NewFriendRepository(friendsstorage.WithFriendRepositoryDB(db))

	userService := usersevice.NewUserService(usersevice.WithUserServiceUserStorage(userRepository))
	postsService := postservice.NewPostService(postservice.WithPostServicePostStorage(postsRepository))
	commentService := commentservice.NewCommentService(commentservice.WithCommentServiceCommentStorage(commentRepository))
	likesService := likesservice.NewLikesService(likesservice.WithLikesServiceLikesStorage(likesRepository))
	friendService := friendsservice.NewFriendService(friendsservice.WithFriendServiceFriendStorage(friendsRepository))

	usersHandler := userhttphandler.NewHandler(userService)
	postsHandler := posthttphandler.NewHandler(postsService)
	commentsHandler := commenthttphandler.NewHandler(commentService)
	likesHandler := likeshttphandler.NewHandler(likesService)
	friendsHandler := friendshttphandler.NewHandler(friendService)

	apiserver.handlers = append(apiserver.handlers, usersHandler, postsHandler, commentsHandler, likesHandler, friendsHandler)

	for _, h := range apiserver.handlers {
		h.AddRoutes(apiserver.app)
	}
}
