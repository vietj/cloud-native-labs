package com.redhat.cloudnative.model;

public class ShoppingCartItem implements java.io.Serializable {
    static final long serialVersionUID = 1L;

   private double price;
	private int quantity;
	private double promoSavings;
	private Product product;
	
	public ShoppingCartItem() {
		
	}
	
	public double getPrice() {
		return price;
	}

	public void setPrice(double price) {
		this.price = price;
	}

	public Product getProduct() {
		return product;
	}

	public void setProduct(Product product) {
		this.product = product;
	}

	public int getQuantity() {
		return quantity;
	}

	public void setQuantity(int quantity) {
		this.quantity = quantity;
	}

	public double getPromoSavings() {
		return promoSavings;
	}

	public void setPromoSavings(double promoSavings) {
		this.promoSavings = promoSavings;
	}

	@Override
	public String toString() {
		return "ShoppingCartItem [price=" + price + ", quantity=" + quantity
				+ ", promoSavings=" + promoSavings + ", product=" + product
				+ "]";
	}
}
