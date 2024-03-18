package service

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/Falokut/admin_casts_service/internal/models"
)

type PersonsChecker interface {
	IsPersonsExists(ctx context.Context, ids []int32) (exists bool, notExistsIDs []int32, err error)
}

type MoviesChecker interface {
	IsMovieExists(ctx context.Context, id int32) (exists bool, err error)
}

type existanceChecker struct {
	moviesCheck  MoviesChecker
	personsCheck PersonsChecker
}

func NewExistanceChecker(moviesCheck MoviesChecker, personsCheck PersonsChecker) *existanceChecker {
	return &existanceChecker{
		moviesCheck:  moviesCheck,
		personsCheck: personsCheck,
	}
}

func convertIntSliceIntoString(nums []int32) string {
	var str = make([]string, len(nums))
	for i, num := range nums {
		str[i] = fmt.Sprint(num)
	}
	return strings.Join(str, ",")
}

func (c *existanceChecker) CheckPersons(ctx context.Context, persons []models.Person) error {
	var ids = make([]int32, len(persons))
	for i := range persons {
		ids[i] = persons[i].ID
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		exists, notFound, err := c.personsCheck.IsPersonsExists(ctx, ids)
		if err != nil {
			return err
		}

		if !exists {
			return models.Errorf(models.InvalidArgument,
				fmt.Sprintf("persons with ids %s not found", convertIntSliceIntoString(notFound)))
		}
	}

	return nil
}

func (c *existanceChecker) CheckMovie(ctx context.Context, id int32) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		exists, err := c.moviesCheck.IsMovieExists(ctx, id)
		if err != nil {
			return err
		} else if !exists {
			return models.Errorf(models.InvalidArgument, "movie with id %d not found", id)
		}
	}

	return nil
}

func (c *existanceChecker) CheckExistance(ctx context.Context, persons []models.Person, movieID int32) error {
	var wg sync.WaitGroup
	reqCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var errCh = make(chan error, 2)

	wg.Add(2)
	go func() {
		defer wg.Done()
		errCh <- c.CheckMovie(reqCtx, movieID)
	}()
	go func() {
		defer wg.Done()
		errCh <- c.CheckPersons(reqCtx, persons)
	}()
	go func() {
		wg.Wait()
		close(errCh)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err, open := <-errCh:
			if !open {
				return nil
			} else if err != nil {
				return err
			}
		}
	}
}
