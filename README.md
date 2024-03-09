# Các Mẫu Design Pattern Và Implement Trong Go

## 1.Option Function Pattern 

### Mô Tả:
- Khi bạn có 1 service lớn và phức tạp cần được khởi tạo, nó sẽ cần rất nhiều tham số để khởi tạo mà chúng ta không muốn quan tâm đến thứ tự truyền vào hoặc không truyền (nếu không truyền thì hàm New sẽ lấy giá trị mặc định)
- Sử dụng 1 function khởi tạo với đâ đủ các tham số chứ không cần phải get rùi set từng tham số, dưới dạng là những function với With
- Thường được sử dụng cho các thư viện trong Go: Grpc, opentelemetry.

### Sử Dụng:
- Khi báo 1 function tên là Option truyền vào con trỏ của complex service. Có tác dụng thay đổi giá trị của các Field trong complex service
- Trong function NewService, truyền vào 1 mảng Option và trả về giá trị của complex service đã tạo.
- Sau đó duyệt qua mảng option và truyền con trỏ service vào `options[i](&service)` tương đương với việc gọi các function With... tương ứng, nó sẽ làm 1 công việc gì đó với service và sau cùng sẽ trả về 1 service với đầy đủ giá trị cần