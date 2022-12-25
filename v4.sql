CREATE TABLE "brands"(
    "id" INTEGER NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NULL,
    "email" VARCHAR(255) NULL,
    "slug" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NULL,
    "logo" VARCHAR(255) NULL,
    "banner" VARCHAR(255) NULL,
    "vendor_id" INTEGER NOT NULL,
    "prefix" VARCHAR(255) NULL,
    "status" VARCHAR(255) NOT NULL,
    "type" VARCHAR(255) NULL,
    "rating" INTEGER NULL,
    "hub_id" INTEGER NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NOT NULL
);
CREATE INDEX "brands_id_index" ON
    "brands"("id");
CREATE INDEX "brands_vendor_id_index" ON
    "brands"("vendor_id");
CREATE INDEX "brands_hub_id_index" ON
    "brands"("hub_id");
ALTER TABLE
    "brands" ADD PRIMARY KEY("id");
CREATE TABLE "products"(
    "id" INTEGER NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "slug" VARCHAR(255) NOT NULL,
    "icon" VARCHAR(255) NULL,
    "images" JSON NULL,
    "details" VARCHAR(255) NULL,
    "price" INTEGER NOT NULL,
    "type" VARCHAR(255) NULL,
    "status" VARCHAR(255) NOT NULL,
    "brand_id" INTEGER NULL,
    "category_id" INTEGER NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "products_brand_id_index" ON
    "products"("brand_id");
CREATE INDEX "products_category_id_index" ON
    "products"("category_id");
CREATE INDEX "products_id_index" ON
    "products"("id");
ALTER TABLE
    "products" ADD PRIMARY KEY("id");
ALTER TABLE
    "products" ADD CONSTRAINT "products_slug_unique" UNIQUE("slug");
CREATE TABLE "order_ratings"(
    "id" INTEGER NOT NULL,
    "order_id" INTEGER NOT NULL,
    "rating" INTEGER NOT NULL,
    "review" VARCHAR(255) NULL,
    "customer_id" INTEGER NOT NULL,
    "pickup_id" INTEGER NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "order_ratings_order_id_index" ON
    "order_ratings"("order_id");
CREATE INDEX "order_ratings_pickup_id_index" ON
    "order_ratings"("pickup_id");
CREATE INDEX "order_ratings_customer_id_index" ON
    "order_ratings"("customer_id");
CREATE INDEX "order_ratings_rating_index" ON
    "order_ratings"("rating");
CREATE INDEX "order_ratings_id_index" ON
    "order_ratings"("id");
ALTER TABLE
    "order_ratings" ADD PRIMARY KEY("id");
CREATE TABLE "orders"(
    "id" INTEGER NOT NULL,
    "delivered_to" INTEGER NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "platform" VARCHAR(255) NULL,
    "dispatch_time" VARCHAR(255) NULL,
    "rider_note" VARCHAR(255) NULL,
    "confirmed_at" DATE NULL,
    "assigned_rider" INTEGER NULL,
    "hub_id" INTEGER NULL,
    "charges" INTEGER NULL,
    "edt" INTEGER NULL,
    "completed_at" DATE NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL,
    "payment_method" VARCHAR(255) NOT NULL,
    "payment_status" VARCHAR(255) NOT NULL
);
CREATE INDEX "orders_delivered_to_index" ON
    "orders"("delivered_to");
CREATE INDEX "orders_hub_id_index" ON
    "orders"("hub_id");
CREATE INDEX "orders_assigned_rider_index" ON
    "orders"("assigned_rider");
CREATE INDEX "orders_id_index" ON
    "orders"("id");
ALTER TABLE
    "orders" ADD PRIMARY KEY("id");
CREATE TABLE "order_items"(
    "id" INTEGER NOT NULL,
    "product_id" INTEGER NOT NULL,
    "quantity" INTEGER NOT NULL,
    "sale_unit" INTEGER NOT NULL,
    "total" INTEGER NOT NULL,
    "discount" INTEGER NOT NULL,
    "pickup_id" INTEGER NULL,
    "order_id" INTEGER NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "order_items_pickup_id_index" ON
    "order_items"("pickup_id");
CREATE INDEX "order_items_order_id_index" ON
    "order_items"("order_id");
CREATE INDEX "order_items_product_id_index" ON
    "order_items"("product_id");
CREATE INDEX "order_items_id_index" ON
    "order_items"("id");
ALTER TABLE
    "order_items" ADD PRIMARY KEY("id");
CREATE TABLE "Pickup"(
    "id" INTEGER NOT NULL,
    "brand_id" INTEGER NOT NULL,
    "order_id" INTEGER NOT NULL,
    "note" VARCHAR(255) NOT NULL,
    "total" INTEGER NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "pickup_brand_id_index" ON
    "Pickup"("brand_id");
CREATE INDEX "pickup_order_id_index" ON
    "Pickup"("order_id");
CREATE INDEX "pickup_status_index" ON
    "Pickup"("status");
CREATE INDEX "pickup_id_index" ON
    "Pickup"("id");
ALTER TABLE
    "Pickup" ADD PRIMARY KEY("id");
CREATE TABLE "delivery_address"(
    "id" INTEGER NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NULL,
    "area" VARCHAR(255) NULL,
    "lat" INTEGER NULL,
    "lng" INTEGER NULL,
    "order_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "delivery_address_order_id_index" ON
    "delivery_address"("order_id");
CREATE INDEX "delivery_address_user_id_index" ON
    "delivery_address"("user_id");
CREATE INDEX "delivery_address_phone_index" ON
    "delivery_address"("phone");
CREATE INDEX "delivery_address_id_index" ON
    "delivery_address"("id");
ALTER TABLE
    "delivery_address" ADD PRIMARY KEY("id");
CREATE TABLE "assigned_rider"(
    "id" INTEGER NOT NULL,
    "order_id" INTEGER NOT NULL,
    "rider_id" INTEGER NOT NULL,
    "movements" JSON NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "assigned_rider_order_id_index" ON
    "assigned_rider"("order_id");
CREATE INDEX "assigned_rider_rider_id_index" ON
    "assigned_rider"("rider_id");
CREATE INDEX "assigned_rider_id_index" ON
    "assigned_rider"("id");
ALTER TABLE
    "assigned_rider" ADD PRIMARY KEY("id");
CREATE TABLE "payment_logs"(
    "id" INTEGER NOT NULL,
    "order_id" INTEGER NOT NULL,
    "method" VARCHAR(255) NOT NULL,
    "trx_id" VARCHAR(255) NOT NULL,
    "amount" INTEGER NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NOT NULL
);
CREATE INDEX "payment_logs_trx_id_index" ON
    "payment_logs"("trx_id");
CREATE INDEX "payment_logs_order_id_index" ON
    "payment_logs"("order_id");
CREATE INDEX "payment_logs_id_index" ON
    "payment_logs"("id");
ALTER TABLE
    "payment_logs" ADD PRIMARY KEY("id");
CREATE TABLE "riders"(
    "id" INTEGER NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "image" VARCHAR(255) NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "riders_phone_index" ON
    "riders"("phone");
CREATE INDEX "riders_id_index" ON
    "riders"("id");
ALTER TABLE
    "riders" ADD PRIMARY KEY("id");
ALTER TABLE
    "riders" ADD CONSTRAINT "riders_phone_unique" UNIQUE("phone");
CREATE TABLE "customers"(
    "id" INTEGER NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "created_at" DATE NOT NULL,
    "deleted_at" DATE NULL,
    "updated_at" DATE NULL
);
CREATE INDEX "customers_email_index" ON
    "customers"("email");
CREATE INDEX "customers_username_index" ON
    "customers"("username");
CREATE INDEX "customers_phone_index" ON
    "customers"("phone");
CREATE INDEX "customers_id_index" ON
    "customers"("id");
ALTER TABLE
    "customers" ADD PRIMARY KEY("id");
ALTER TABLE
    "customers" ADD CONSTRAINT "customers_phone_unique" UNIQUE("phone");
ALTER TABLE
    "customers" ADD CONSTRAINT "customers_username_unique" UNIQUE("username");
ALTER TABLE
    "customers" ADD CONSTRAINT "customers_email_unique" UNIQUE("email");
CREATE TABLE "order_timeline"(
    "id" INTEGER NOT NULL,
    "order_id" INTEGER NOT NULL,
    "message" INTEGER NOT NULL,
    "action_type" INTEGER NOT NULL,
    "referance_id" INTEGER NULL,
    "created_at" DATE NOT NULL,
    "deleted_at" DATE NULL,
    "updated_at" DATE NULL
);
CREATE INDEX "order_timeline_order_id_index" ON
    "order_timeline"("order_id");
CREATE INDEX "order_timeline_referance_id_index" ON
    "order_timeline"("referance_id");
CREATE INDEX "order_timeline_id_index" ON
    "order_timeline"("id");
ALTER TABLE
    "order_timeline" ADD PRIMARY KEY("id");
CREATE TABLE "cart_items"(
    "id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "product_id" INTEGER NOT NULL,
    "quantity" INTEGER NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "cart_items_user_id_index" ON
    "cart_items"("user_id");
CREATE INDEX "cart_items_product_id_index" ON
    "cart_items"("product_id");
CREATE INDEX "cart_items_id_index" ON
    "cart_items"("id");
ALTER TABLE
    "cart_items" ADD PRIMARY KEY("id");
CREATE TABLE "hub"(
    "id" INTEGER NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NULL,
    "deleted_at" DATE NULL
);
CREATE INDEX "hub_id_index" ON
    "hub"("id");
ALTER TABLE
    "hub" ADD PRIMARY KEY("id");
CREATE TABLE "OrderCharges"(
    "id" INTEGER NOT NULL,
    "total" INTEGER NOT NULL,
    "total_discount" INTEGER NULL,
    "service_charge" INTEGER NULL,
    "delivery_charge" INTEGER NOT NULL,
    "item_discount" INTEGER NULL,
    "promo_discount" INTEGER NULL,
    "order_id" INTEGER NOT NULL,
    "voucher" JSON NULL,
    "created_at" DATE NOT NULL,
    "updated_at" DATE NOT NULL,
    "deleted_at" DATE NOT NULL
);
CREATE INDEX "ordercharges_order_id_index" ON
    "OrderCharges"("order_id");
ALTER TABLE
    "OrderCharges" ADD PRIMARY KEY("id");
ALTER TABLE
    "order_ratings" ADD CONSTRAINT "order_ratings_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE
    "order_ratings" ADD CONSTRAINT "order_ratings_customer_id_foreign" FOREIGN KEY("customer_id") REFERENCES "customers"("id");
ALTER TABLE
    "order_ratings" ADD CONSTRAINT "order_ratings_pickup_id_foreign" FOREIGN KEY("pickup_id") REFERENCES "Pickup"("id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_hub_id_foreign" FOREIGN KEY("hub_id") REFERENCES "hub"("id");
ALTER TABLE
    "order_items" ADD CONSTRAINT "order_items_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "products"("id");
ALTER TABLE
    "order_items" ADD CONSTRAINT "order_items_pickup_id_foreign" FOREIGN KEY("pickup_id") REFERENCES "Pickup"("id");
ALTER TABLE
    "order_items" ADD CONSTRAINT "order_items_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE
    "Pickup" ADD CONSTRAINT "pickup_brand_id_foreign" FOREIGN KEY("brand_id") REFERENCES "brands"("id");
ALTER TABLE
    "Pickup" ADD CONSTRAINT "pickup_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE
    "delivery_address" ADD CONSTRAINT "delivery_address_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "customers"("id");
ALTER TABLE
    "assigned_rider" ADD CONSTRAINT "assigned_rider_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE
    "assigned_rider" ADD CONSTRAINT "assigned_rider_rider_id_foreign" FOREIGN KEY("rider_id") REFERENCES "riders"("id");
ALTER TABLE
    "payment_logs" ADD CONSTRAINT "payment_logs_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE
    "order_timeline" ADD CONSTRAINT "order_timeline_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");
ALTER TABLE
    "cart_items" ADD CONSTRAINT "cart_items_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "customers"("id");
ALTER TABLE
    "cart_items" ADD CONSTRAINT "cart_items_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "products"("id");
ALTER TABLE
    "OrderCharges" ADD CONSTRAINT "ordercharges_order_id_foreign" FOREIGN KEY("order_id") REFERENCES "orders"("id");