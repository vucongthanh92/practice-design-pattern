# Các Mẫu Design Pattern Và Implement Trong Go

## A. Creational Pattern

### 1. Option Function

#### Mô Tả:
- Khi bạn có 1 service lớn và phức tạp cần được khởi tạo, nó sẽ cần rất nhiều tham số để khởi tạo mà chúng ta không muốn quan tâm đến thứ tự truyền vào hoặc không truyền (nếu không truyền thì hàm New sẽ lấy giá trị mặc định)
- Sử dụng 1 function khởi tạo với đâ đủ các tham số chứ không cần phải get rùi set từng tham số, dưới dạng là những function với With
- Thường được sử dụng cho các thư viện trong Go: Grpc, opentelemetry.

#### Sử Dụng:
- Khi báo 1 function tên là Option truyền vào con trỏ của complex service. Có tác dụng thay đổi giá trị của các Field trong complex service. `type Option func(*complexSvc)`
- Trong function NewService, truyền vào 1 mảng Option và trả về giá trị của complex service đã tạo.

```` 
func NewService(opts ...Option) complexSvc {
  service := complexSvc{
  name:      "Service",
  logger:    StdLogger{},
  notifier:  SMSNotifier{},
  DataLayer: MongoDB{},
  uploader:  AwsS3Uploader{},
  }

  for i := range opts {
  opts[i](&service)
  }

  return service
 }
````

- Sau đó duyệt qua mảng option và truyền con trỏ service vào `options[i](&service)` tương đương với việc gọi các function With... tương ứng, nó sẽ làm 1 công việc gì đó với service và sau cùng sẽ trả về 1 service với đầy đủ giá trị cần
````
func WithMysqlDB() Option {
	return func(s *complexSvc) {
		s.DataLayer = MysqlDB{}
	}
}
````

## C. Behavioral Pattern
- Đây là những design pattern liên quan đến hành vị của các object
### 2. Iterator
#### Mô Tả:
- Đây là design pattern giúp cho developer duyệt qua các thành phần trong 1 collection mà không cần biết bên dưới nó là dạng thể hiện nào như là: linked-list, array, tree.
- Dùng 1 interface để làm tham số đầu vào chung, sau đó implement các method để xử l riêng cho từng loại dữ liệu.
#### Sử Dụng:
- Tạo ra một interface iterator với 2 method là getNext(lấy phần tử tiếp theo) và hasMore(xác định xem có phần tử kế tiếp) để duyệt qua tất cả element ccua3tham số truyền vào.
````
type TransferIterator interface {
	Next() TransferInterface
	HasNext() bool
}
````
- Tạo ra các concrete iterator để xử lý riêng cho từng loại dữ liệu(array, linked-list, tree...) và implement vào interface iterator đã tạo ở trên.
````
type ArrayIterator struct {
	index int
	arr   []TransferInterface
}

func (a *ArrayIterator) HasNext() bool {
	return len(a.arr) > 0 && a.index < len(a.arr)
}

func (a *ArrayIterator) Next() TransferInterface {
	item := a.arr[a.index]
	a.index++
	return item
}
````
- Ở đây, ta cũng cần function để tạo ra một concrete iterator của kiểu dữ liệu và trả về. Làm tương tự cho linked-list và tree nếu có.
````
func NewImplementArray(arr []TransferInterface) TransferIterator {
	return &ArrayIterator{index: 0, arr: arr}
}
````
- Tạo ra một function để duyệt qua các element của đối tượng và thực thi các process mong muốn. Dùng HasNext để làm điều kiện dừng cho vòng lặp for khi đã duyệt hết các phần tử và dùng `Next()` để thực hiện tiếp process với phần tử tiếp theo.
````
func Deposit(iterator TransferIterator, amount float64) {
	for iterator.HasNext() {
		iterator.Next().Receive(amount)
	}
}
````
- Với kiểu dữ liệu là linked-list và tree thì chúng ta có thêm 1 struct để lưu giá trị con trỏ của phần tử hiện tại, và dùng object này để thực hiện các method `HasNext()` và `Next()` khi mà duyệt qua các element.
````
type LinkedPointer struct {
	node *LinkedNode
}

type TreePointer struct {
	node *TreeNode
}
````
- Tạo các data mockup simple để test thử 
````
var ArrayMockupData = []TransferInterface{
	Profile{name: "Peter", balance: 0},
	Profile{name: "Mary", balance: 0},
	Profile{name: "Tom", balance: 0},
	Profile{name: "Harry", balance: 0},
}
var LinkedListMockupData = &LinkedNode{
	val: Profile{name: "Peter", balance: 10},
	next: &LinkedNode{
		val: Profile{name: "Mary", balance: 10},
		next: &LinkedNode{
			val: Profile{name: "Tom", balance: 10},
			next: &LinkedNode{
				val:  Profile{name: "Harry", balance: 10},
				next: nil,
			},
		},
	},
}
var TreeMockupData = &TreeNode{
	val: Profile{name: "Peter", balance: 25},
	children: []TreeNode{
		{
			val: Profile{name: "Tom", balance: 25},
			children: []TreeNode{
				{val: Profile{name: "Mary", balance: 25}},
				{val: Profile{name: "Vincent", balance: 25}},
				{val: Profile{name: "Vicky", balance: 25}},
			},
		},
		{
			val: Profile{name: "bob", balance: 25},
			children: []TreeNode{
				{val: Profile{name: "Alice", balance: 25}},
			},
		},
	},
}
````
- Cuối cùng là thực thi, tạo ra một đối tượng iterator tương ứng ví dụ như linked-list thì sẽ dùng function `NewImplementLinkedList()`. Sau đó chúng ta chỉ cần gọi hàm muốn thực thi là `Deposit()`
````
fmt.Println("\n****** a->b->c Linked-list iterator ******")
iterator = internal.NewImplementLinkedList(internal.LinkedListMockupData)
internal.Deposit(iterator, amount)
````