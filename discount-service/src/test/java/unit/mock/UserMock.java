package unit.mock;

import java.time.LocalDate;

import com.fredoliveira.discountcalculator.domain.User;

import static java.time.Month.FEBRUARY;

public class UserMock {

  public static User getOne() {
    return User.builder()
      .id("a50a93f9-748a-476e-ac26-74314b84e304")
      .dateOfBirth(LocalDate.of(1988, FEBRUARY, 19))
      .firstName("John")
      .lastName("Rambo")
      .build();
  }

}
