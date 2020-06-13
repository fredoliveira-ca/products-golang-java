package com.fredoliveira.discountcalculator;

import java.io.IOException;

import com.fredoliveira.discountcalculator.app.grpc.GrpcServer;

public class Application {
	public static void main(String[] args) throws IOException, InterruptedException {
		GrpcServer.start();
	}
}
