package com.fredoliveira.discountcalculator.app.utility;

import java.time.Clock;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.ZoneId;

import lombok.experimental.UtilityClass;

@UtilityClass
public class DeLoreanMachine {

  private final ZoneId zoneId = ZoneId.systemDefault();
  private Clock clock = Clock.systemDefaultZone();

  public LocalDate getToday() {
    return LocalDate.now(getClock());
  }

  public void travelToPresent() {
    clock = Clock.systemDefaultZone();
  }

  public void travelTo(LocalDate date) {
    var dateTime = LocalDateTime.of(date, LocalTime.MIDNIGHT);
    clock = Clock.fixed(dateTime.atZone(zoneId).toInstant(), zoneId);
  }

  private Clock getClock() {
    return clock ;
  }
}
