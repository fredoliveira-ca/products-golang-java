package integration.com.fredoliveira.discountcalculator.app.grpc;

import java.io.IOException;

import com.fredoliveira.discountcalculator.app.grpc.product.FetchProductGrpc;
import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.app.service.DiscountService;
import com.fredoliveira.discountcalculator.app.service.DiscountStrategy;
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
import org.junit.jupiter.api.Test;

import static com.fredoliveira.discountcalculator.domain.Promotion.BIRTHDAY;
import static net.bytebuddy.implementation.bytecode.constant.FloatConstant.ZERO;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

public class GrpcServerTest {

  @Rule
  public final GrpcCleanupRule grpcCleanup = new GrpcCleanupRule();

  @Test
  void shouldGetFivePercentDiscountWhenIsUsersBirthday() throws IOException {
    FetchUserGrpc userGrpc = mock(FetchUserGrpc.class);
    FetchProductGrpc productGrpc = mock(FetchProductGrpc.class);
    DiscountStrategy strategy = new DiscountStrategy();

    final var product = ProductMock.getOne();
    when(productGrpc.fetchBy(any())).thenReturn(product);
    when(userGrpc.fetchBy(any())).thenReturn(UserMock.getOneWhoIsBirthdayToday());

    DiscountServiceGrpc.DiscountServiceBlockingStub blockingStub = registerGrpcMock(
      userGrpc, productGrpc, strategy);

    final var calculate = blockingStub.calculate(DiscountRequest.newBuilder().build());

    assertNotEquals(ZERO, calculate.getPct());
    assertEquals(BIRTHDAY.getDiscount().floatValue(), calculate.getPct());
    assertEquals(
      MoneyUtils.getDiscountValue(product.getPriceInCents(), BIRTHDAY.getDiscount()),
      calculate.getValueInCents());
  }

  @Test
  void shouldGetZeroDiscountWhenCouldNotFetchUser() throws IOException {
    FetchUserGrpc userGrpc = new FetchUserGrpc();
    FetchProductGrpc productGrpc = mock(FetchProductGrpc.class);
    DiscountStrategy strategy = new DiscountStrategy();
    when(productGrpc.fetchBy(any())).thenReturn(ProductMock.getOne());

    DiscountServiceGrpc.DiscountServiceBlockingStub blockingStub = registerGrpcMock(
      userGrpc, productGrpc, strategy);

    final var calculate = blockingStub.calculate(DiscountRequest.newBuilder().build());

    assertEquals(0, calculate.getPct());
    assertEquals(0, calculate.getValueInCents());
  }

  @Test
  void shouldAbortCalculateWhenCouldNotFetchProduct() throws IOException {
    FetchUserGrpc userGrpc = new FetchUserGrpc();
    FetchProductGrpc productGrpc = new FetchProductGrpc();
    DiscountStrategy strategy = new DiscountStrategy();

    DiscountServiceGrpc.DiscountServiceBlockingStub blockingStub = registerGrpcMock(
      userGrpc, productGrpc, strategy);

    assertThrows(
      StatusRuntimeException.class,
      () -> blockingStub.calculate(DiscountRequest.newBuilder().build()),
      "io.grpc.StatusRuntimeException: ProductGrpcServer is down!"
    );
  }

  private DiscountServiceGrpc.DiscountServiceBlockingStub registerGrpcMock(FetchUserGrpc userGrpc, FetchProductGrpc productGrpc, DiscountStrategy strategy) throws IOException {
    String serverName = InProcessServerBuilder.generateName();

    grpcCleanup.register(
      InProcessServerBuilder
        .forName(serverName)
        .directExecutor()
        .addService(new DiscountService(userGrpc, productGrpc, strategy))
        .build()
        .start()
    );

    return DiscountServiceGrpc
      .newBlockingStub(grpcCleanup.register(
        InProcessChannelBuilder.forName(serverName)
          .directExecutor()
          .build()));
  }
}
