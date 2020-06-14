package com.fredoliveira.discountcalculator.domain;

import java.time.LocalDate;

import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

@Builder
@Getter
@Setter
public class User {

  private final String id;
  private final String firstName;
  private final String lastName;
  private LocalDate dateOfBirth;

}
