package com.fredoliveira.discountcalculator.app.utility;

import java.time.LocalDate;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static java.time.LocalDate.now;
import static java.time.Month.JUNE;
import static java.time.Month.SEPTEMBER;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;

@DisplayName("Runs all tests for utility layer of time machine")
public class DeLoreanMachineTest {

  @Test
  @DisplayName("should return today")
  void getToday() {
    final var today = DeLoreanMachine.getToday();

    assertEquals(now(), today);
  }

  @Test
  @DisplayName("should travel to September, 2020 and go back today")
  void returnToToday() {
    LocalDate date = LocalDate.of(2020, SEPTEMBER, 15);

    DeLoreanMachine.travelTo(date);
    assertEquals(date, DeLoreanMachine.getToday());

    DeLoreanMachine.travelToPresent();
    assertEquals(LocalDate.now(), DeLoreanMachine.getToday());

  }

  @Test
  @DisplayName("should travel to June, 1988")
  void calculateWithoutDiscount() {
    LocalDate travelTo = LocalDate.of(1988, JUNE, 13);

    DeLoreanMachine.travelTo(travelTo);
    assertEquals(travelTo, DeLoreanMachine.getToday());

    DeLoreanMachine.travelToPresent();
    assertNotEquals(travelTo, DeLoreanMachine.getToday());
  }

}
