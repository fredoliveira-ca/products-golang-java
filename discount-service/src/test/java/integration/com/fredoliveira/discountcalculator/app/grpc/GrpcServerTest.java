package integration.com.fredoliveira.discountcalculator.app.grpc;

import java.io.IOException;

import com.fredoliveira.discountcalculator.app.grpc.product.FetchProductGrpc;
import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.app.service.DiscountService;
import com.fredoliveira.discountcalculator.app.service.DiscountStrategy;
import com.fredoliveira.discountcalculator.app.utility.DeLoreanMachine;
import com.fredoliveira.discountcalculator.app.utility.MoneyUtils;
import integration.mock.ProductMock;
import integration.mock.UserMock;
import io.grpc.StatusRuntimeException;
import io.grpc.inprocess.InProcessChannelBuilder;
import io.grpc.inprocess.InProcessServerBuilder;
import io.grpc.product.DiscountRequest;
import io.grpc.product.DiscountServiceGrpc;
import io.grpc.testing.GrpcCleanupRule;
import org.junit.Rule;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static com.fredoliveira.discountcalculator.domain.Discount.LIMIT_DISCOUNT;
import static com.fredoliveira.discountcalculator.domain.Promotion.BIRTHDAY;
import static com.fredoliveira.discountcalculator.domain.Promotion.BLACK_FRIDAY;
import static net.bytebuddy.implementation.bytecode.constant.FloatConstant.ZERO;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

@DisplayName("Runs all tests for gRPC layer of discount calculator")
public class GrpcServerTest {

  @Rule
  public final GrpcCleanupRule grpcTest = new GrpcCleanupRule();
  public final FetchUserGrpc userGrpc = mock(FetchUserGrpc.class);
  public final FetchProductGrpc productGrpc = mock(FetchProductGrpc.class);
  public final DiscountStrategy strategy = new DiscountStrategy();
  public final DiscountServiceGrpc.DiscountServiceBlockingStub blockingStub;

  public GrpcServerTest() throws IOException {
    blockingStub = registerGrpcMock(userGrpc, productGrpc, strategy);
  }

  @BeforeEach
  void setup() {
    DeLoreanMachine.travelToPresent();
  }

  @Test
  void shouldGetPercentageDiscountLimitWhenIsBlackFridayAndUsersBrithday() {
    final var product = ProductMock.getOne();
    when(productGrpc.fetchBy(any())).thenReturn(product);
    when(userGrpc.fetchBy(any())).thenReturn(UserMock.getOneWhoBirthdayIsOnBlackFriday());

    DeLoreanMachine.travelTo(BLACK_FRIDAY.getPromotionDate());
    final var calculate = blockingStub.calculate(DiscountRequest.newBuilder().build());

    assertNotEquals(ZERO, calculate.getPct());
    assertEquals(LIMIT_DISCOUNT.floatValue(), calculate.getPct());
    assertEquals(
      MoneyUtils.getDiscountValue(product.getPriceInCents(), LIMIT_DISCOUNT),
      calculate.getValueInCents());
  }

  @Test
  void shouldGetTenPercentDiscountWhenIsBlackFriday() {
    final var product = ProductMock.getOne();
    when(productGrpc.fetchBy(any())).thenReturn(product);
    when(userGrpc.fetchBy(any())).thenReturn(UserMock.getOne());

    DeLoreanMachine.travelTo(BLACK_FRIDAY.getPromotionDate());
    final var calculate = blockingStub.calculate(DiscountRequest.newBuilder().build());

    assertNotEquals(ZERO, calculate.getPct());
    assertEquals(BLACK_FRIDAY.getDiscount().floatValue(), calculate.getPct());
    assertEquals(
      MoneyUtils.getDiscountValue(product.getPriceInCents(), BLACK_FRIDAY.getDiscount()),
      calculate.getValueInCents());
  }

  @Test
  void shouldGetFivePercentDiscountWhenIsUsersBirthday() {
    final var product = ProductMock.getOne();
    when(productGrpc.fetchBy(any())).thenReturn(product);
    when(userGrpc.fetchBy(any())).thenReturn(UserMock.getOneWhoBirthdayIsToday());

    final var calculate = blockingStub.calculate(DiscountRequest.newBuilder().build());

    assertNotEquals(ZERO, calculate.getPct());
    assertEquals(BIRTHDAY.getDiscount().floatValue(), calculate.getPct());
    assertEquals(
      MoneyUtils.getDiscountValue(product.getPriceInCents(), BIRTHDAY.getDiscount()),
      calculate.getValueInCents());
  }

  @Test
  void shouldGetZeroDiscountWhenCouldNotFetchUser() {
    when(productGrpc.fetchBy(any())).thenReturn(ProductMock.getOne());

    final var calculate = blockingStub.calculate(DiscountRequest.newBuilder().build());

    assertEquals(0, calculate.getPct());
    assertEquals(0, calculate.getValueInCents());
  }

  @Test
  void shouldAbortCalculateWhenCouldNotFetchProduct() {
    assertThrows(
      StatusRuntimeException.class,
      () -> blockingStub.calculate(DiscountRequest.newBuilder().build()),
      "io.grpc.StatusRuntimeException: ProductGrpcServer is down!"
    );
  }

  private DiscountServiceGrpc.DiscountServiceBlockingStub registerGrpcMock(
    FetchUserGrpc userGrpc,
    FetchProductGrpc productGrpc,
    DiscountStrategy strategy
  ) throws IOException {
    String serverName = InProcessServerBuilder.generateName();

    grpcTest.register(
      InProcessServerBuilder
        .forName(serverName)
        .directExecutor()
        .addService(new DiscountService(userGrpc, productGrpc, strategy))
        .build()
        .start()
    );

    return DiscountServiceGrpc
      .newBlockingStub(grpcTest.register(
        InProcessChannelBuilder.forName(serverName)
          .directExecutor()
          .build()));
  }
}
