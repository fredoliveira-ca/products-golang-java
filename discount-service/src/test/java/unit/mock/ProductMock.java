package unit.mock;

import com.fredoliveira.discountcalculator.domain.Product;

public class ProductMock {

  public static Product getOne() {
    return Product.builder()
      .id("e2303619-27ff-4661-80c6-ffcd70d04909")
      .priceInCents(100L)
      .title("Ball")
      .description("Rounded")
      .build();
  }

  public static Product getOneFree() {
    return Product.builder()
      .id("e2303619-27ff-4661-80c6-ffcd70d04909")
      .priceInCents(0L)
      .title("Ball")
      .description("Rounded")
      .build();
  }

}
