package main
 import(
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
 )

 type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed  bool  `json:"completed"`
 }
 var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record viseo", Completed: false},
 }

 func getTodos(context *gin.Context){
    context.IndentedJSON(http.StatusOK, todos )
 }
func addTodo(context *gin.Context){
	var newTodo todo 

	if err := context.BindJSON(&newTodo); err  != nil {
		return
	}
	todos =append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
def toggoleTodStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil{
		context.IndentedJSON(http.StatusnotFound, gin.H{"message":"todo not found"})
		return
}
todo.Completed = !todo.completed

context.IndentedJSON(http.StatusOK, todo)
}
func getTodoId(id string) (*todo, error){
	for i, t:= range todos{
		if t.ID == id{
			return &todos[i], nil
		}
	}
	 return nil, errors.New("todos not found")
}
func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil{
		context.IndentedJSON(http.StatusnotFound, gin.H{"message":"todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}
 func main(){
      router :=gin.Default()
	  router.GET("/todos",getTodos )
	  router.PATCH("/todos/:id",getTodos )
	  router.POST("/todos", addTodo)
      router.Run("localhost:9090")
 }