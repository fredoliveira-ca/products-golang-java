package com.fredoliveira.discountcalculator.app.service;

import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.app.utility.DeLoreanMachine;
import com.fredoliveira.discountcalculator.testetet.mock.ProductMock;
import com.fredoliveira.discountcalculator.testetet.mock.UserMock;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static com.fredoliveira.discountcalculator.domain.Discount.LIMIT_DISCOUNT;
import static com.fredoliveira.discountcalculator.domain.Promotion.BIRTHDAY;
import static com.fredoliveira.discountcalculator.domain.Promotion.BLACK_FRIDAY;
import static java.math.BigDecimal.ZERO;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

@DisplayName("Runs all tests for service layer of discount calculator")
public class DiscountCalculatorTest {

  private final FetchUserGrpc userGrpc = mock(FetchUserGrpc.class);
  private DiscountService service;

  @BeforeEach
  void beforeEach() {
    DeLoreanMachine.travelToPresent();
    DiscountStrategy strategy = new DiscountStrategy();
    service = new DiscountService(userGrpc, strategy);
  }

  @Test
  @DisplayName("should calculate and retun no discount")
  void calculateWithoutDiscount() {
    final var product = ProductMock.getProductMock();
    final var user = UserMock.getUserMock();
    when(userGrpc.fetchBy(any())).thenReturn(user);

    final var discount = service.calculateDiscount(product.getPriceInCents(), user.getId());

    assertEquals(ZERO, discount.getPercentage());
  }

  @Test
  @DisplayName("should calculate and retun a black friday discount")
  void calculateBlackFridayDiscount() {
    final var product = ProductMock.getProductMock();
    final var user = UserMock.getUserMock();
    when(userGrpc.fetchBy(any())).thenReturn(user);

    DeLoreanMachine.travelTo(BLACK_FRIDAY.getPromotionDate());

    final var discount = service.calculateDiscount(product.getPriceInCents(), user.getId());

    assertEquals(BLACK_FRIDAY.getDiscount(), discount.getPercentage());
    assertNotEquals(ZERO, discount.getPercentage());
  }

  @Test
  @DisplayName("should calculate and retun a birthday discount")
  void calculateBirthdayDiscount() {
    final var product = ProductMock.getProductMock();
    final var user = UserMock.getUserMock();
    when(userGrpc.fetchBy(any())).thenReturn(user);
    DeLoreanMachine.travelTo(user.getDateOfBirth());

    final var discount = service.calculateDiscount(product.getPriceInCents(), user.getId());

    assertEquals(BIRTHDAY.getDiscount(), discount.getPercentage());
  }

  @Test
  @DisplayName("should calculate sum birthday and friday discounts and retun the discount limit")
  void calculateAllDiscount() {
    final var product = ProductMock.getProductMock();
    final var user = UserMock.getUserMock();
    user.setDateOfBirth(BLACK_FRIDAY.getPromotionDate());

    when(userGrpc.fetchBy(any())).thenReturn(user);

    DeLoreanMachine.travelTo(BLACK_FRIDAY.getPromotionDate());

    final var discount = service.calculateDiscount(product.getPriceInCents(), user.getId());

    assertEquals(LIMIT_DISCOUNT, discount.getPercentage());
  }

}
