package com.fredoliveira.discountcalculator.app.grpc.user;

import java.time.LocalDate;

import com.fredoliveira.discountcalculator.domain.User;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.user.UserRequest;
import io.grpc.user.UserResponse;
import io.grpc.user.UserServiceGrpc;

public class FetchUserGrpc {
  public static final String ADDRESS = "localhost";
  public static final int PORT = 50053;

  public User fetchBy(String userId) {
    ManagedChannel channel = ManagedChannelBuilder.forAddress(ADDRESS, PORT)
      .usePlaintext()
      .build();

    UserServiceGrpc.UserServiceBlockingStub stub
      = UserServiceGrpc.newBlockingStub(channel);

    UserResponse response = stub.fetchOne(
      UserRequest.newBuilder()
        .setUserId(userId)
        .build()
    );

    channel.shutdown();

    final var dateOfBirth = response.getUser().getDateOfBirth();
    return User.builder()
      .id(userId)
      .dateOfBirth(
        LocalDate.of(
          dateOfBirth.getYear(),
          dateOfBirth.getMonth(),
          dateOfBirth.getDay()))
      .build();
  }
}
