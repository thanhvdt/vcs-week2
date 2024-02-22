package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type elasticService struct {
	ec *elasticsearch.Client
}

//func (es *elasticService) searchCustomer(c *gin.Context) {
//	res, err := es.ec.Search().
//		Index("index_name").
//		Request(&search.Request{
//			Query: &types.Query{
//				Match: map[string]types.MatchQuery{
//					"name": {Query: "Foo"},
//				},
//			},
//		}).Do(context.Background())
//}
