package com.fredoliveira.discountcalculator.app.service;

import com.fredoliveira.discountcalculator.data.BirthdayDiscount;
import com.fredoliveira.discountcalculator.data.BlackFridayDiscount;
import com.fredoliveira.discountcalculator.domain.Promotion;

public class DiscountStrategy {

    public static DiscountFinder of(Promotion promotion, String userId) {
        return switch (promotion) {
            case BIRTHDAY -> new BirthdayDiscount(userId);
            case BLACK_FRIDAY -> new BlackFridayDiscount();
        };
    }

}
