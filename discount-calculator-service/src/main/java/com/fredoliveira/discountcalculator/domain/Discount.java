package com.fredoliveira.discountcalculator.domain;

import java.math.BigDecimal;

import lombok.Builder;
import lombok.Getter;

@Getter
@Builder
public class Discount {

  private final BigDecimal percentage;
  private final Long valueInCents;

}
