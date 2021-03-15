CREATE TABLE headphones (
    id serial PRIMARY KEY,
    name varchar (50) NOT NULL,
    image_url varchar (200) NOT NULL,
    description varchar (400) NOT NULL,
    type integer NOT NULL,
    price float,
    extra_info varchar (400)
);

INSERT INTO headphones VALUES (100, 'Beats EP', 'https://hnsfpau.imgix.net/5/images/detailed/106/ML9D2PA-A.4.jpg?fit=fill&bg=0FFF&w=1500&h=844&auto=format,compress',
'This is a test description for a Beats EP headphone', 10, 500, 'Available in three colors - red, blue, white');
INSERT INTO headphones VALUES (101, 'Air Pods', 'https://images-na.ssl-images-amazon.com/images/I/71zny7BTRlL._AC_SL1500_.jpg',
'This is a test description for a Air Pods', 10, 300, 'Available in white');