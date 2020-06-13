package com.fredoliveira.discountcalculator.app.service;

import java.time.LocalDate;

import com.fredoliveira.discountcalculator.app.grpc.product.FetchProductGrpc;
import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.app.service.DiscountCalculator;
import com.fredoliveira.discountcalculator.domain.Product;
import com.fredoliveira.discountcalculator.domain.User;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static java.math.BigDecimal.ZERO;
import static java.time.Month.FEBRUARY;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

@DisplayName("Runs all tests for service layer of discount calculator")
public class DiscountCalculatorTest {

  private final FetchProductGrpc productGrpc = mock(FetchProductGrpc.class);
  private final FetchUserGrpc userGrpc = mock(FetchUserGrpc.class);
  private DiscountCalculator service;

  @BeforeEach void beforeEach() {
    service = new DiscountCalculator();
  }

  @Test
  @DisplayName("should calculate and retun no discount")
  void calculateWithoutDiscount() {
    final var product = Product.builder()
      .id("234").priceInCents(1000L)
      .build();

    final var user = User.builder()
      .id("123").dateOfBirth(LocalDate.of(1988, FEBRUARY, 19))
      .build();

    when(productGrpc.fetchBy(any())).thenReturn(product);
    when(userGrpc.fetchBy(any())).thenReturn(user);

    final var discount = service.calculateDiscount(product.getPriceInCents(), user.getId());

    assertEquals(ZERO, discount.getPercentage());
  }

  @Test
  @DisplayName("should calculate and retun a black friday discount")
  void calculateBlackFridayDiscount() {
    final var product = Product.builder()
      .id("234").priceInCents(1000L)
      .build();

    final var user = User.builder()
      .id("123").dateOfBirth(LocalDate.of(1988, FEBRUARY, 19))
      .build();

    final var discount = service.calculateDiscount(product.getPriceInCents(), user.getId());

    assertEquals(ZERO, discount.getPercentage());
  }

}
