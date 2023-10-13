package repository_test

/*
func TestSelfSave(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	var goodm connection.Goodm
	ctx := context.Background()

	mt.Run("SUCCESS - TestFindOnePrivate", func(mt *mtest.T) {

		goodm.ConnectMock(mt.Client).UseDatabase("mock_db", &ctx)

		for _, us := range SelfSaveValues {
			user := repository.NewGoodmDoc[*UserMock](
				&UserMock{
					GoodmCollection: repository.GoodmCollection[*UserMock]{
						Id: primitive.NewObjectID(),
					},
					Name:      us.Name,
					Username:  us.Username,
					Latitude:  us.Latitude,
					Longitude: us.Longitude,
					MailCount: us.MailCount,
				},
			)

			mt.AddMockResponses(mtest.CreateSuccessResponse())
			mt.AddMockResponses(mtest.CreateCursorResponse(0, "mock_db.users", mtest.FirstBatch, bson.D{
				{Key: "_id", Value: user.Document.Id},
				{Key: "Name", Value: user.Document.Name},
				{Key: "Username", Value: user.Document.Username},
				{Key: "Latitude", Value: user.Document.Latitude},
				{Key: "Longitude", Value: user.Document.Longitude},
				{Key: "MailCount", Value: user.Document.MailCount},
			}))

			savedUser, err := user.Save()

			assert.NotNil(t, *savedUser)
			assert.Nil(t, err)
			assert.Equal(t, *savedUser, *user.Document)
		}

	})

}
*/
