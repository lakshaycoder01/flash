project by flash monk


apis :

 POST /v2/flash_monk/add_product  // Add products to product table 

`json: {
    "brand" "nike",
    "name": "vtshirt",
    "price": 1500.00,
    "quantity": 50
}
    `

 GET API Call  /v2/flash_monk/view_products


POST /v2/flash_monk/buy_product

`json: {
    "customer_id": 150,
    "product_id": 130,
    "quantity": 20,
}
    `

Here quantity is a pointer as input , can be available or not , if not available default quantity is 1 otherwise quantity

/v2/flash_monk/cancel_product

`json: {
    "customer_id": 150,
    "product_id": 130,
    "quantity": 20,
}
    `


GET Call /v2/flash_monk/fetch_customer_orders/:customerID

will fetch all the orders for customers which are in status ('CONDIRMED', 'DELIVERED')


/v2/flash_monk/search_products

`json: {
   "brand_name": "nike",
   "product_name": "tshirt",
   "price": 150.0
}
    `


schema:

CREATE TABLE `customer` (
    `id` int(20) AUTO_INCREMENT, 
    `name` varchar(200) NOT NULL,
    `email` varchar(200) NOT NULL,
    `created_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `i_customer_id` (`id`)`,
    KEY `i_customer_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `products` (
    `id` int(20) AUTO_INCREMENT, 
    `name` varchar(200) NOT NULL,
    `price` float NOT NULL,
    `brand` varchar(200) NOT NULL,
    `quantity` int NOT NULL,
    `created_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `i_products_id` (`id`)`
    KEY `i_products_name` (`name`)
    KEY `i_products_brand` (`brand`),
    KEY `i_products_price` (`price`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `customer_product` (
    `id` int(20) AUTO_INCREMENT, 
    `customer_id` int(20) NOT NULL,
    `product_id` int(20) NOT NULL,
    `brand` varchar(200) NOT NULL,
    `product_name` varchar(200) NOT NULL,
    `price` float NOT NULL,
    `quantity` int NOT NULL,
    `created_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `i_customer_product_id (`id`)`
    KEY `i_customer_product_customer_id` (`customer_id`)
    KEY `i_customer_product_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


All these values are pointers request object which can be present in payload or not, if present result will filter based on what values is present