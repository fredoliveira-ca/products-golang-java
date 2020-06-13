package com.fredoliveira.discountcalculator.app.service;

import com.fredoliveira.discountcalculator.app.grpc.product.FetchProductGrpc;
import com.fredoliveira.discountcalculator.domain.Discount;
import io.grpc.product.DiscountRequest;
import io.grpc.product.DiscountResponse;
import io.grpc.product.DiscountServiceGrpc;
import io.grpc.stub.StreamObserver;

public class DiscountService extends DiscountServiceGrpc.DiscountServiceImplBase {

  @Override public void calculate(DiscountRequest request, StreamObserver<DiscountResponse> responseObserver) {
    final var product = new FetchProductGrpc().fetchBy(request.getProductId());
    final var discount = new DiscountCalculator().calculateDiscount(product.getPriceInCents(), request.getUserId());
    respond(responseObserver, discount);
  }

  private void respond(StreamObserver<DiscountResponse> responseObserver, Discount discount) {
    DiscountResponse response = DiscountResponse.newBuilder()
        .setPct(discount.getPercentage().floatValue())
        .setValueInCents(discount.getValueInCents())
        .build();

    responseObserver.onNext(response);
    responseObserver.onCompleted();
  }
}
