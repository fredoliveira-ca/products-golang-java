package com.fredoliveira.discountcalculator.app.service;

import java.math.BigDecimal;

import com.fredoliveira.discountcalculator.app.grpc.product.FetchProductGrpc;
import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.domain.Discount;
import com.fredoliveira.discountcalculator.domain.Promotion;
import io.grpc.product.DiscountRequest;
import io.grpc.product.DiscountResponse;
import io.grpc.product.DiscountServiceGrpc;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;

import static com.fredoliveira.discountcalculator.domain.Discount.DEFAULT_DISCOUNT;
import static com.fredoliveira.discountcalculator.domain.Discount.LIMIT_DISCOUNT;
import static java.util.Arrays.stream;

@RequiredArgsConstructor
public class DiscountService extends DiscountServiceGrpc.DiscountServiceImplBase {

  private final FetchUserGrpc userGrpc;
  private final DiscountStrategy strategy;

  @Override
  public void calculate(DiscountRequest request, StreamObserver<DiscountResponse> responseObserver) {
    final var product = new FetchProductGrpc().fetchBy(request.getProductId());
    final var discount = calculateDiscount(product.getPriceInCents(), request.getUserId());
    respond(responseObserver, discount);
  }

  public Discount calculateDiscount(Long priceInCents, String userId) {
    final var totalDiscount = stream(Promotion.values())
        .map(promotion -> strategy.of(promotion, userId, userGrpc).getDiscount())
        .reduce(BigDecimal.ZERO, BigDecimal::add);

    var discount = DEFAULT_DISCOUNT;
    if (totalDiscount.compareTo(DEFAULT_DISCOUNT) > 0) {
      discount = LIMIT_DISCOUNT.min(totalDiscount);
    }

    return Discount.builder()
        .build()
        .calculate(priceInCents, discount);
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
