package integration.mock;

import java.time.LocalDate;

import com.fredoliveira.discountcalculator.domain.User;

import static com.fredoliveira.discountcalculator.domain.Promotion.BLACK_FRIDAY;
import static java.time.Month.FEBRUARY;

public class UserMock {

  public static final String USER_ID = "a50a93f9-748a-476e-ac26-74314b84e304";

  public static User getOne() {
    return User.builder()
      .id(USER_ID)
      .dateOfBirth(LocalDate.of(1988, FEBRUARY, 19))
      .firstName("John")
      .lastName("Rambo")
      .build();
  }

  public static User getOneWhoBirthdayIsToday() {
    return User.builder()
      .id(USER_ID)
      .dateOfBirth(LocalDate.now())
      .firstName("John")
      .lastName("Rambo")
      .build();
  }

  public static User getOneWhoBirthdayIsOnBlackFriday() {
    return User.builder()
      .id(USER_ID)
      .dateOfBirth(BLACK_FRIDAY.getPromotionDate())
      .firstName("John")
      .lastName("Rambo")
      .build();
  }
}
