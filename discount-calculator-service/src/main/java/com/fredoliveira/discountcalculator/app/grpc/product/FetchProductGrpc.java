package com.fredoliveira.discountcalculator.app.grpc.product;

import com.fredoliveira.discountcalculator.domain.Product;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.product.ProductPriceRequest;
import io.grpc.product.ProductPriceResponse;
import io.grpc.product.ProductPriceServiceGrpc;

public class FetchProductGrpc {
  public static final String ADDRESS = "localhost";
  public static final int PORT = 50051;

  public Product fetchBy(String productId) {
    ManagedChannel channel = ManagedChannelBuilder.forAddress(ADDRESS, PORT)
      .usePlaintext()
      .build();

    ProductPriceServiceGrpc.ProductPriceServiceBlockingStub stub
      = ProductPriceServiceGrpc.newBlockingStub(channel);

    ProductPriceResponse response = stub.fetchOne(
      ProductPriceRequest.newBuilder()
        .setProductId(productId)
        .build()
    );

    channel.shutdown();

    return Product.builder()
      .id(productId)
      .priceInCents(response.getValueInCents())
      .build();
  }
}
