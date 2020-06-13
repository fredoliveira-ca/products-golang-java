package com.fredoliveira.discountcalculator;

import com.fredoliveira.discountcalculator.app.grpc.GrpcServer;

public class Application {
  public static void main(String[] args) {
    GrpcServer.start();
  }
}
