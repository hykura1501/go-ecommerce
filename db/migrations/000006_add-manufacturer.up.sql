-- Xóa toàn bộ dữ liệu hiện có trong bảng manufacturer
DELETE FROM manufacturer;

-- Chèn các hãng sản xuất vào bảng manufacturer
INSERT INTO manufacturer (manufacturer_name) VALUES
  ('Logitech'),
  ('Samsung'),
  ('Sony'),
  ('Apple'),
  ('Asus'),
  ('Lenovo');
