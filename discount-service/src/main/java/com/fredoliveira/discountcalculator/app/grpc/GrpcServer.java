package com.fredoliveira.discountcalculator.app.grpc;

import com.fredoliveira.discountcalculator.app.grpc.product.FetchProductGrpc;
import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.app.service.DiscountStrategy;
import com.fredoliveira.discountcalculator.app.service.DiscountService;
import io.grpc.ServerBuilder;
import lombok.AccessLevel;
import lombok.AllArgsConstructor;

@AllArgsConstructor(access = AccessLevel.PRIVATE)
public final class GrpcServer {
  private static final java.util.logging.Logger log = java.util.logging.Logger.getLogger(GrpcServer.class.getName());

  public static void start() {
    final FetchUserGrpc userGrpc = new FetchUserGrpc();
    final DiscountStrategy strategy = new DiscountStrategy();
    final FetchProductGrpc productGrpc = new FetchProductGrpc();

    try {
      io.grpc.Server server = ServerBuilder
          .forPort(50052)
          .addService(new DiscountService(userGrpc, productGrpc, strategy))
          .build();

      server.start();
      server.awaitTermination();
      log.info("Server is up!");
    } catch (Exception e) {
      log.severe("Something is wrong here:" + e);
    }
  }

}
