package drivers

import "github.com/jcasado94/nats-points/mongo/entity"

func (md *MongoDriver) GetArticlesUrl(countryName, newspaper string) (string, error) {
	return md.countryService.GetArticlesUrl(countryName, newspaper)
}

// func (md *MongoDriver) DeleteAllCountryArticles(countryName string) ([]bson.ObjectId, error) {
// 	return md.countryService.DeleteAllArticles(countryName)
// }

// func (md *MongoDriver) AddAllArticles(countryName string, modelArticles []model.ArticleModel) error {
// 	return md.countryService.AddAllArticles(countryName, modelArticles)
// }

func (md *MongoDriver) GetAllCountryArticlesTagged(countryName string) (map[string][]entity.Article, error) {
	country, err := md.countryService.GetCountry(countryName)
	if err != nil {
		return nil, err
	}
	res := make(map[string][]entity.Article)
	for tag, ids := range country.Articles {
		res[tag] = md.GetArticles(ids)
	}
	return res, nil
}

func (md *MongoDriver) GetAllCountryArticles(countryName string) ([]entity.Article, error) {
	ids, err := md.countryService.GetAllArticlesIDs(countryName)
	if err != nil {
		return nil, err
	}
	return md.GetArticles(ids), nil
}

// func (md *MongoDriver) GetAllCountryResultArticles(countryName string) ([]entity.Article, error) {
// 	taggedArticles, err := md.GetAllCountryResultArticlesTagged(countryName)
// 	if err != nil {
// 		return entity.SortableArticles{}, err
// 	}
// 	return taggedArticles.MergeTags().GetElements(), nil
// }

// func (md *MongoDriver) GetAllCountryResultArticlesTagged(countryName string) (entity.ResultArticlesTagged, error) {
// 	articles, err := md.GetAllCountryArticles(countryName)
// 	if err != nil {
// 		return entity.ResultArticlesTagged{}, err
// 	}
// 	mappedArticles, err := md.articlesService.GetAllArticlesMapped()
// 	if err != nil {
// 		return entity.ResultArticlesTagged{}, err
// 	}
// 	res := entity.NewResultArticles()
// 	for _, artId := range articles.Environment {
// 		modelArt := mappedArticles[artId]
// 		art := model.ArticleFromModel(&modelArt)
// 		res.Environment.Add(&art)
// 	}
// 	res.Environment = res.Environment.GetElements()
// 	for _, artId := range articles.Politics {
// 		modelArt := mappedArticles[artId]
// 		art := model.ArticleFromModel(&modelArt)
// 		res.Politics.Add(&art)
// 	}
// 	res.Politics = res.Politics.GetElements()
// 	for _, artId := range articles.Society {
// 		modelArt := mappedArticles[artId]
// 		art := model.ArticleFromModel(&modelArt)
// 		res.Society.Add(&art)
// 	}
// 	res.Society = res.Society.GetElements()
// 	for _, artId := range articles.Sport {
// 		modelArt := mappedArticles[artId]
// 		art := model.ArticleFromModel(&modelArt)
// 		res.Sports.Add(&art)
// 	}
// 	res.Sports = res.Sports.GetElements()
// 	for _, artId := range articles.Business {
// 		modelArt := mappedArticles[artId]
// 		art := model.ArticleFromModel(&modelArt)
// 		res.Business.Add(&art)
// 	}
// 	res.Business = res.Business.GetElements()
// 	for _, artId := range articles.Culture {
// 		modelArt := mappedArticles[artId]
// 		art := model.ArticleFromModel(&modelArt)
// 		res.Culture.Add(&art)
// 	}
// 	res.Culture = res.Culture.GetElements()

// 	return res, nil
// // }

func (md *MongoDriver) GetCountryInformation(countryName string) (entity.Information, error) {
	country, err := md.countryService.GetCountry(countryName)
	if err != nil {
		return entity.Information{}, err
	}
	return country.Info, nil
}
