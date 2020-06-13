package com.fredoliveira.discountcalculator.domain;

import java.time.LocalDate;

import lombok.Builder;
import lombok.Getter;

@Builder
@Getter
public class User {

  private final String id;
  private final String firstName;
  private final String lastName;
  private final LocalDate dateOfBirth;

}
