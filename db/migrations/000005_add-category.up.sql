-- Xóa hết dữ liệu cũ
DELETE FROM category;

-- Insert các category cha
INSERT INTO category (name, description, thumbnail)
VALUES
('PC Game Headsets', 'High-quality headsets designed for immersive PC gaming experiences.', 'https://media.istockphoto.com/id/1403828056/vector/circuit-board-blue-technology-background.jpg?s=612x612&w=0&k=20&c=FI1jC45r-VAe9heJFR_SGOGosbEi-zNm3B-SgDJJ25c='),
('Computers & Tablets', 'Latest computers and tablets for all your computing needs.', 'https://media.tapchitaichinh.vn/w1480/images/upload//2022/12/27/tai-chinh-cong-nghe-2.jpg'),
('PlayStation 5 Headsets', 'Top-notch headsets for an enhanced PlayStation 5 gaming experience.', 'https://media.tapchitaichinh.vn/w1480/images/upload/tranhuyentrang/01152020/fintech_financial_technology_icons_circuit_board_thinkstock_664731514_3x2-100736056-large.jpg'),
('PlayStation 4 Headsets', 'High-quality headsets for PlayStation 4 gamers.', 'https://fbu.edu.vn/wp-content/uploads/2020/02/nganh-cong-nghe-thong-tin.jpg'),
('Computer Headsets', 'Comfortable and high-fidelity headsets for computer use.', 'https://www.pace.edu.vn/uploads/news/ImageContent/2015/09/cong-nghe-thong-tin-trong-san-xuat.png'),
('Computer Keyboards, Mice & Accessories', 'A wide range of keyboards, mice, and accessories for your computer.', 'https://media.tapchitaichinh.vn/w1480/images/upload//2022/12/27/tai-chinh-cong-nghe-2.jpg'),
('Computer Keyboards', 'Durable and ergonomic keyboards for all types of computer users.', 'https://media.tapchitaichinh.vn/w1480/images/upload/hoxuantruong/08302022/09.JPG'),
('Computer Keyboard & Mouse Combos', 'Convenient keyboard and mouse combos for seamless computing.', 'https://media.doanhnhantrevietnam.vn/files/content/2023/08/12/tien-kts-0851.jpg'),
('Mac Games & Accessories', 'Games and accessories specifically designed for Mac users.', 'https://www.flashfly.net/wp/wp-content/uploads/2022/09/14-vs-16-inch-mbp-m2-pro-and-max-feature-1.jpeg'),
('Mac Gaming Keyboards', 'High-performance gaming keyboards for Mac users.', 'https://media.istockphoto.com/id/1363326235/vi/anh/flat-lay-c%E1%BB%A7a-c%C3%A1c-s%E1%BA%A3n-ph%E1%BA%A9m-t%C3%A1o-kh%C3%A1c-nhau-tr%C3%AAn-n%E1%BB%81n-m%C3%A0u-x%C3%A1m.jpg?s=612x612&w=0&k=20&c=LjLfDJR6Fkhi3_txTqxyxUie77klD66HlwJSZ_W9ZNY='),
('Asus Gaming Keyboards', 'Top-tier gaming keyboards from Asus for an enhanced gaming experience.', 'https://fbu.edu.vn/wp-content/uploads/2020/02/nganh-cong-nghe-thong-tin.jpg'),
('Mac Gaming Mice', 'Precision gaming mice designed for Mac users.', 'https://media.tapchitaichinh.vn/w1480/images/upload//2022/12/27/tai-chinh-cong-nghe-2.jpg'),
('Gaming Laptops', 'High-performance laptops built for gaming.', 'https://media.tapchitaichinh.vn/w1480/images/upload//2022/12/27/tai-chinh-cong-nghe-2.jpg'),
('PC Gaming Accessories', 'Essential accessories to enhance your PC gaming setup.', 'https://fbu.edu.vn/wp-content/uploads/2020/02/nganh-cong-nghe-thong-tin.jpg'),
('RGB Gaming Keyboards', 'Vibrant RGB gaming keyboards for a visually stunning gaming experience.', 'https://media.doanhnhantrevietnam.vn/files/content/2023/08/12/tien-kts-0851.jpg'),
('High-Performance Mice', 'Precision mice designed for high-performance computing and gaming.', 'https://www.flashfly.net/wp/wp-content/uploads/2022/09/14-vs-16-inch-mbp-m2-pro-and-max-feature-1.jpeg'),
('Custom Build PCs', 'Custom-built PCs tailored to your specific needs and preferences.', 'https://media.istockphoto.com/id/1363326235/vi/anh/flat-lay-c%E1%BB%A7a-c%C3%A1c-s%E1%BA%A3n-ph%E1%BA%A9m-t%C3%A1o-kh%C3%A1c-nhau-tr%C3%AAn-n%E1%BB%81n-m%C3%A0u-x%C3%A1m.jpg?s=612x612&w=0&k=20&c=LjLfDJR6Fkhi3_txTqxyxUie77klD66HlwJSZ_W9ZNY=');

-- Update super_category_id (sử dụng subquery để lấy đúng id)
UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Game Headsets')
WHERE name = 'Computers & Tablets';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'Computers & Tablets')
WHERE name = 'PlayStation 5 Headsets';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Game Headsets')
WHERE name = 'PlayStation 4 Headsets';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Game Headsets')
WHERE name = 'Computer Headsets';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'Computer Headsets')
WHERE name = 'Computer Keyboards, Mice & Accessories';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Game Headsets')
WHERE name = 'Computer Keyboards';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'Computers & Tablets')
WHERE name = 'Computer Keyboard & Mouse Combos';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'Computer Headsets')
WHERE name = 'Mac Games & Accessories';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'Mac Gaming Keyboards')
WHERE name = 'Asus Gaming Keyboards';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Gaming Accessories')
WHERE name = 'Computers & Tablets';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'Gaming Laptops')
WHERE name = 'Custom Build PCs';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Gaming Accessories')
WHERE name = 'RGB Gaming Keyboards';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PC Gaming Accessories')
WHERE name = 'High-Performance Mice';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PlayStation 4 Headsets')
WHERE name = 'Mac Gaming Mice';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'Computers & Tablets')
WHERE name = 'Gaming Laptops';

UPDATE category
SET super_category_id = (SELECT category_id FROM category WHERE name = 'PlayStation 4 Headsets')
WHERE name = 'Computer Headsets';
