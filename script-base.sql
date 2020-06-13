CREATE DATABASE productdb;

CREATE TABLE product (
    product_id text PRIMARY KEY,
    price_in_cents integer,
    title text,
    description text
);

CREATE TABLE "user" (
    user_id text PRIMARY KEY,
    first_name text,
    last_name text,
    date_of_birth date NOT NULL
);

INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('9325817d-f543-4718-9621-6d42d93d73f4', 2899, 'Adult Toronto Raptors Fanatics Branded Patriotic Face Covering 3-Pack', 'This product is not a medical device. Cover up while you represent your favorite team in this Toronto Raptors face covering.');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('e4d3f34c-d316-4b3b-aa26-84d776e32a41', 2899, 'Men''s Chicago Bulls New Era Black Rugged Canvas 9TWENTY Snap 2 Adjustable Hat', 'Add another eye-catching layer of Chicago Bulls flair to your outfit when you choose this riveting Rugged Canvas Snap 2 adjustable hat from New Era.');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('74bb9928-78f6-4370-93c2-c4f8b0ab1227', 4999, 'Fire 7 Tablet, 7” display', '16 GB, 7” touchscreen, 1024 x 600 resolution at 171 ppi, SD video playback, with IPS (in-plane switching) technology and advanced polarizing filter');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('4bf6bf60-45d2-4a25-a205-3fd978b88673', 2499, 'Echo Dot (3rd Gen) - Smart speaker with Alexa - Charcoal', 'With a new speaker and design, Echo Dot is a voice-controlled smart speaker with Alexa, perfect for any room. Just ask for music, news, information, and more. You can also call almost anyone and control compatible smart home devices with your voice.');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('4aad7605-bb1c-4082-9af7-f56f2b4289ae', 3199, 'Men''s Toronto Raptors Red #1 Dad T-Shirt', 'Get your dad this Toronto Raptors #1 Dad tee to show him how much you appreciate him! It features bold graphics for a look he''ll love.');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('ac615ff3-209e-4520-b136-61b07f0b62f5', 499, 'Find Me (Inland Empire Book 1)', 'Convicted serial killer Benjamin Fisher has finally offered to lead San Bernardino detective Daniel Ellis to the isolated graves of his victims.');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('471e4d5f-c614-4425-b45c-4ae4278dbe1f', 1759, 'LIFX A19 Mini Wi-Fi Smart Led Light Bulb (Latest Generation), Dimmable, Warm White, No Hub Required, Works with Amazon Alexa, Apple HomeKit, Google Assistant and Microsoft Cortana', '650 lumens; 50W comparable');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('5882c679-6806-45dd-b79f-8789562a0848', 3299, 'Men''s Toronto Raptors Nike Black 2019/20 City Edition Swingman Shorts', 'You''re an avid Toronto Raptors fan and love to flaunt it. Show the Toronto Raptors your support by grabbing these 2019/20 City Edition Swingman Shorts from Nike.');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('1faa5d7f-8a39-4425-9804-7acc369d66ad', 2499, 'Remo HD-8510-00 Fiberskyn Frame Drum, 10"', 'Constructed with an Acousticon drum shell and a pre-tuned Fiberskyn drumhead, these frame drums produce excellent projection and enhanced low-pitched fundamentals typically found with traditional wood shell drums, making them ideal for professional and Recreational applications.');
INSERT INTO product(
            product_id, price_in_cents, title, description)
    VALUES ('5d140bd1-5892-4f9e-80c6-e37d4b9f621a', 1214, 'Vic Firth Signature Series -- Steve Jordan Drumsticks (SJOR)', 'Steve Jordan is well known as a multi-instrumentalist, musical director, producer and a writer of exceptional quality. His most recent projects involve performing with the John Mayer Trio and Sonny Rollins, while in the past he has collaborated with countless super stars including Billy Joel, Bruce Springsteen, Alicia Keys, Stevie Nicks, Luther Vandross, Neil Young, Sheryl Crow, B. B. King, Bob Dylan, James Taylor and David Sanborn. His new Signature stick is light and long for great touch and sound around the drums and cymbals. In hickory. L = 16 ½ ; Dia. = . 525');

INSERT INTO "user"(
            user_id, first_name, last_name, date_of_birth)
    VALUES ('41597637-8c33-409f-a869-a2090e87ec78', 'John', 'Generated', '1988-02-19');
INSERT INTO "user"(
            user_id, first_name, last_name, date_of_birth)
    VALUES ('4a07fb31-d908-411b-949e-6ae3effbe60b', 'Anna', 'Provided', '2000-02-27');
INSERT INTO "user"(
            user_id, first_name, last_name, date_of_birth)
    VALUES ('9e2a8a33-941b-4e31-a282-3dba00b4302c', 'Anna', 'Provided', CURRENT_DATE);