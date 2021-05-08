package message

import (
	"testing"
)

func Test_PostgresRepo_Get(t *testing.T) {
	// var a, _ = userFixture.CreateUserFixture(nil, "")
	// fmt.Print(a)

	// var inMemoryRepoTestDeps = repoTestDependencies{
	// 	users: []*User{
	// 		userFromMemoryMap(*memstore.NewUserFixture(1)),
	// 		userFromMemoryMap(*memstore.NewUserFixture(2)),
	// 	},
	// }
	// testRepositoryGet(
	// 	t,
	// 	inMemoryRepoTestDeps,
	// 	func() Repository {
	// 		store := memstore.NewStore()

	// 		for _, user := range inMemoryRepoTestDeps.users {
	// 			testUser := memstore.NewUserFixture(user.ID)
	// 			store.GetUsers().Put(testUser.ID, *testUser)
	// 		}

	// 		return NewInMemoryRepo(store)
	// 	},
	// )
}
