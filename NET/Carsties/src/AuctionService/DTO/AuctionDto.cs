using System;

namespace AuctionService.DTO;

public class AuctionDto
{
  public Guid Id { get; set; }
  public int ReservePrice { get; set; }
  public string Seller { get; set; }
  public int SoldAmount { get; set; }
  public int CurrentHigh { get; set; }
  public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
  public DateTime UpdatedAt { get; set; } = DateTime.UtcNow;
  public DateTime AuctionEnd { get; set; }
  public string Status { get; set; }

  public string Make { get; set; }
  public string Model { get; set; }
  public int Year { get; set; }
  public string Color { get; set; }
  public int Mileage { get; set; }
  public string ImageUrl { get; set; }
}
