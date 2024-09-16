package resolver

import (
	"errors"
	"github.com/Lux00000/post-and-comments/internal/models"
	"log"
	"sync"
)

type CommentsObservers struct {
	chans   map[int][]CommentObserver
	counter int
	mu      sync.Mutex
}

type CommentObserver struct {
	ch chan *models.Comment
	id int
}

func NewCommentsObserver() *CommentsObservers {
	return &CommentsObservers{
		chans:   make(map[int][]CommentObserver),
		counter: 0,
		mu:      sync.Mutex{},
	}
}

func (c *CommentsObservers) CreateObserver(postId int) (int, chan *models.Comment, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch := make(chan *models.Comment)
	c.counter++

	c.chans[postId] = append(c.chans[postId], CommentObserver{ch: ch, id: c.counter})

	return c.counter, ch, nil
}

func (c *CommentsObservers) DeleteObserver(postId, chanId int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	obs, ok := c.chans[postId]
	if !ok {
		return errors.New("observer not found")
	}

	for i, observer := range obs {
		if observer.id == chanId {
			close(observer.ch)
			c.chans[postId] = append(obs[:i], obs[i+1:]...)
			return nil
		}
	}

	return errors.New("observer not found")
}

func (c *CommentsObservers) NotifyObservers(postId int, comment models.Comment) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	obs, ok := c.chans[postId]
	if !ok {
		return errors.New("observers not found")
	}

	for _, observer := range obs {
		select {
		case observer.ch <- &comment:
		default:
			log.Println("NotifyObservers: channel is full, skipping")
		}
	}

	return nil
}
