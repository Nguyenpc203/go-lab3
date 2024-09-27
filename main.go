package main

import (
	"fmt"
	"math"
)

// 1. Quản lý thông tin nhân viên
type NhanVien struct {
	ten     string
	tuoi    int
	chucVu  string
}

func NhanVienCoTuoiCaoNhat(nhanViens []NhanVien) NhanVien {
	nhanVienLonTuoiNhat := nhanViens[0]
	for _, nhanVien := range nhanViens {
		if nhanVien.tuoi > nhanVienLonTuoiNhat.tuoi {
			nhanVienLonTuoiNhat = nhanVien
		}
	}
	return nhanVienLonTuoiNhat
}

// 2. Tính chu vi và diện tích hình tròn
type HinhTron struct {
	banKinh float64
}

func (h HinhTron) ChuVi() float64 {
	return 2 * math.Pi * h.banKinh
}

func (h HinhTron) DienTich() float64 {
	return math.Pi * math.Pow(h.banKinh, 2)
}

// Interface Hình học
type HinhDang interface {
	DienTich() float64
	ChuVi() float64
}

type HinhChuNhat struct {
	chieuRong, chieuDai float64
}

func (h HinhChuNhat) DienTich() float64 {
	return h.chieuRong * h.chieuDai
}

func (h HinhChuNhat) ChuVi() float64 {
	return 2 * (h.chieuRong + h.chieuDai)
}

// Interface Phương tiện giao thông
type PhuongTien interface {
	KhoiDong()
	DungLai()
}

type OTo struct {
	mauXe string
}

func (o OTo) KhoiDong() {
	fmt.Printf("Oto %s đã khởi động\n", o.mauXe)
}

func (o OTo) DungLai() {
	fmt.Printf("Oto %s đã dừng lại\n", o.mauXe)
}

type XeMay struct {
	nhanHieu string
}

func (x XeMay) KhoiDong() {
	fmt.Printf("Xe máy %s đã khởi động\n", x.nhanHieu)
}

func (x XeMay) DungLai() {
	fmt.Printf("Xe máy %s đã dừng lại\n", x.nhanHieu)
}

// Con trỏ
func HoanDoi(so1, so2 *int) {
	*so1, *so2 = *so2, *so1
}

func TangGiaTri(so *int) {
	*so++
}

// Đệ quy
func GiaiThua(n int) int {
	if n == 0 {
		return 1
	}
	return n * GiaiThua(n-1)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Defer, Panic, Recover
func ChiaAnToan(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Lỗi:", r)
		}
	}()
	if b == 0 {
		panic("Không thể chia cho 0")
	}
	fmt.Printf("Kết quả: %d\n", a/b)
}

func DeferVaPanic() {
	defer fmt.Println("Tạm biệt!")
	fmt.Println("Xin chào!")
	panic("Đã xảy ra lỗi")
}

// Hàm Variadic
func Tong(nums ...int) int {
	tong := 0
	for _, num := range nums {
		tong += num
	}
	return tong
}

func GiaTriLonNhat(nums ...int) int {
	if len(nums) == 0 {
		panic("Không có số nào")
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	// Quản lý thông tin nhân viên
	fmt.Println("=== Bài tập 1: Quản lý thông tin nhân viên ===")
	nhanViens := []NhanVien{
		{"John", 30, "Quản lý"},
		{"Alice", 25, "Kỹ sư"},
		{"Bob", 40, "Thiết kế"},
	}
	nhanVienLonTuoiNhat := NhanVienCoTuoiCaoNhat(nhanViens)
	fmt.Println("Nhân viên lớn tuổi nhất:", nhanVienLonTuoiNhat.ten)

	// Tính chu vi và diện tích hình tròn
	fmt.Println("\n=== Bài tập 2: Tính chu vi và diện tích hình tròn ===")
	hinhTron := HinhTron{banKinh: 5}
	fmt.Printf("Chu vi: %.2f, Diện tích: %.2f\n", hinhTron.ChuVi(), hinhTron.DienTich())

	// Hình học với interface
	fmt.Println("\n=== Bài tập 3: Hình học với interface ===")
	hinhChuNhat := HinhChuNhat{chieuRong: 4, chieuDai: 6}
	fmt.Printf("Diện tích Hình chữ nhật: %.2f, Chu vi Hình chữ nhật: %.2f\n", hinhChuNhat.DienTich(), hinhChuNhat.ChuVi())

	// Phương tiện giao thông
	fmt.Println("\n=== Bài tập 4: Phương tiện giao thông ===")
	oto := OTo{"Tesla Model 3"}
	xemay := XeMay{"Yamaha"}
	oto.KhoiDong()
	xemay.KhoiDong()
	oto.DungLai()
	xemay.DungLai()

	// Hoán đổi hai số
	fmt.Println("\n=== Bài tập 5: Hoán đổi hai số ===")
	so1, so2 := 3, 5
	HoanDoi(&so1, &so2)
	fmt.Println("Sau khi hoán đổi: so1 =", so1, ", so2 =", so2)

	// Tăng giá trị
	fmt.Println("\n=== Bài tập 6: Tăng giá trị ===")
	so := 10
	TangGiaTri(&so)
	fmt.Println("Sau khi tăng giá trị:", so)

	// Tính giai thừa
	fmt.Println("\n=== Bài tập 7: Tính giai thừa ===")
	n := 5
	fmt.Printf("%d! = %d\n", n, GiaiThua(n))

	// Dãy Fibonacci
	fmt.Println("\n=== Bài tập 8: Dãy Fibonacci ===")
	fibNum := Fibonacci(6)
	fmt.Println("Fibonacci thứ 6:", fibNum)

	// Xử lý lỗi trong tính toán
	fmt.Println("\n=== Bài tập 9: Xử lý lỗi trong tính toán ===")
	ChiaAnToan(10, 0)

	// Tính toán với defer và panic
	// fmt.Println("\n=== Bài tập 10: Tính toán với defer và panic ===")
	// DeferVaPanic()

	// Tính tổng các số
	fmt.Println("\n=== Bài tập 11: Tính tổng các số ===")
	fmt.Println("Tổng:", Tong(1, 2, 3, 4))

	// Tìm giá trị lớn nhất
	fmt.Println("\n=== Bài tập 12: Tìm giá trị lớn nhất ===")
	fmt.Println("Giá trị lớn nhất:", GiaTriLonNhat(3, 5, 7, 2))
}
