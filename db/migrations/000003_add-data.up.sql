-- Chèn dữ liệu vào bảng category
INSERT INTO category (name, description, thumbnail) VALUES
    ('PC Game Headsets', 'High-quality headsets designed for immersive PC gaming experiences.', 'https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c='),
    ('Computers & Tablets', 'Latest computers and tablets for all your computing needs.', 'https://media.tapchitaichinh.vn/w1480/images/upload//2022/12/27/tai-chinh-cong-nghe-2.jpg'),
    ('PlayStation 5 Headsets', 'Top-notch headsets for an enhanced PlayStation 5 gaming experience.', 'https://media.tapchitaichinh.vn/w1480/images/upload/tranhuyentrang/01152020/fintech_financial_technology_icons_circuit_board_thinkstock_664731514_3x2-100736056-large.jpg'),
    ('PlayStation 4 Headsets', 'High-quality headsets for PlayStation 4 gamers.', 'https://fbu.edu.vn/wp-content/uploads/2020/02/nganh-cong-nghe-thong-tin.jpg'),
    ('Computer Headsets', 'Comfortable and high-fidelity headsets for computer use.', 'https://www.pace.edu.vn/uploads/news/ImageContent/2015/09/cong-nghe-thong-tin-trong-san-xuat.png'),
    ('Computer Keyboards, Mice & Accessories', 'A wide range of keyboards, mice, and accessories for your computer.', 'https://media.tapchitaichinh.vn/w1480/images/upload//2022/12/27/tai-chinh-cong-nghe-2.jpg'),
    ('Computer Keyboards', 'Durable and ergonomic keyboards for all types of computer users.', 'https://media.tapchitaichinh.vn/w1480/images/upload/hoxuantruong/08302022/09.JPG');

-- Cập nhật super_category_id cho các danh mục con
UPDATE category SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Game Headsets') WHERE name = 'Computers & Tablets';
UPDATE category SET super_category_id = (SELECT category_id FROM category WHERE name = 'Computers & Tablets') WHERE name = 'PlayStation 5 Headsets';
UPDATE category SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Game Headsets') WHERE name = 'PlayStation 4 Headsets';
UPDATE category SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Game Headsets') WHERE name = 'Computer Headsets';
UPDATE category SET super_category_id = (SELECT category_id FROM category WHERE name = 'Computer Headsets') WHERE name = 'Computer Keyboards, Mice & Accessories';
UPDATE category SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Game Headsets') WHERE name = 'Computer Keyboards';
UPDATE category SET super_category_id = (SELECT category_id FROM category WHERE name = 'Computers & Tablets') WHERE name = 'Computer Keyboard & Mouse Combos';



-- Chèn dữ liệu vào bảng manufacturer
INSERT INTO manufacturer (manufacturer_name) VALUES
    ('Logitech'),
    ('Samsung'),
    ('Sony'),
    ('Apple'),
    ('Asus'),
    ('Lenovo');




INSERT INTO products (product_name, description, price, stock, category_id, manufacturer_id, discount, type, tag) VALUES
('Logitech MK270 Wireless Keyboard And Mouse Combo', 'Logitech MK270 Wireless Keyboard And Mouse Combo For Windows, 2.4 GHz Wireless, Compact Mouse, 8 Multimedia And Shortcut Keys, For PC, Laptop - Black', 27.99, 100, 6, 1, 0, 'keyboard', 'featured'),
('NPET K10V3PRO Gaming Keyboard', 'NPET K10V3PRO Gaming Keyboard, RGB Backlit Keys, Spill-Resistant, Customizable Keys, Dedicated Multi-Media Keys – Black', 12.49, 90, 6, 2, 0, 'keyboard', 'featured'),
('SteelSeries Arctis Nova 1 Multi-System Gaming Headset', 'SteelSeries Arctis Nova 1 Multi-System Gaming Headset — Hi-Fi Drivers — 360° Spatial Audio — Comfort Design — Durable — Ultra Lightweight — Noise-Cancelling Mic — PC, PS5/PS4, Switch, Xbox - Black', 59.99, 80, 1, 3, 0, 'headset', 'new'),
('Logitech K120 Wired Keyboard for Windows', 'Logitech K120 Wired Keyboard for Windows, Plug and Play, Full-Size, Spill-Resistant, Curved Space Bar, Compatible with PC, Laptop - Black', 12.34, 80, 2, 1, 0.10, 'keyboard', 'sale'),
('SteelSeries Apex 3 RGB Gaming Keyboard', 'SteelSeries Apex 3 RGB Gaming Keyboard – 10-Zone RGB Illumination – IP32 Water Resistant – Premium Magnetic Wrist Rest (Whisper Quiet Gaming Switch)', 35.99, 70, 4, 3, 0, 'keyboard', 'sale'),
('Rii RK907 Ultra-Slim Compact USB Wired Keyboard for Mac and PC', 'Rii RK907 Ultra-Slim Compact USB Wired Keyboard for Mac and PC, Windows 10/8 / 7 / Vista/XP (Black) (1PCS)', 14.99, 85, 2, 3, 0, 'keyboard', 'sale'),
('Logitech K270 Wireless Keyboard for Windows', 'Logitech K270 Wireless Keyboard for Windows, 2.4 GHz Wireless, Full-Size, Number Pad, 8 Multimedia Keys, 2-Year Battery Life, Compatible with PC, Laptop, Black', 24.95, 67, 1, 5, 0, 'keyboard', 'featured'),
('Dell Wired Keyboard', 'Dell Wired Keyboard - Black KB216 (580-ADMT)', 15.99, 70, 1, 6, 0, 'keyboard', 'new'),
('AULA F99 Wireless Mechanical Keyboard', 'AULA F99 Wireless Mechanical Keyboard,Tri-Mode BT5.0/2.4GHz/USB-C Hot Swappable Custom Keyboard,Pre-lubed Linear Switches,Gasket Structure,RGB Backlit Gaming Keyboard for PC/Tablet/PS/Xbox', 69.95, 30, 4, 1, 0, 'keyboard', 'featured'),
('Rii RK100+ Multiple Color Rainbow LED Backlit Keyboard', 'Rii RK100+ Multiple Color Rainbow LED Backlit Large Size USB Wired Mechanical Feeling Multimedia PC Gaming Keyboard,Office Keyboard for Working or Primer Gaming,Office Device', 16.99, 45, 8, 5, 0, 'keyboard', 'sale');
