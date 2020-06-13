package com.fredoliveira.discountcalculator.domain;

import lombok.Builder;
import lombok.Getter;

@Getter
@Builder
public class Product {

  private final String id;
  private final String title;
  private final String description;
  private final Long priceInCents;
  private final Discount discount;

}
