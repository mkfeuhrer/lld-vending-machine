## Notes

### Clarify requirements

- Multiple product support (dynamic price and quantity)
- Accept coins and notes of dynamic denomination
- Return change after product is bought
- Track products and quantities
- Handle multiple transactions (concurrency)
- Refill the supply
- Collect money
- Handle error cases like insufficient funds, out of stock

Questions:

- Do we need to manage change? If coins/notes are available?
  - No
- Do we need to maintain purchase history?
  - Yes and active balances
- Multiple product support in one transaction?
  - Yes

### Break down into Entities

- Product
- Payment
- PaymentType
  - Coin
  - Notes
  - Credit card (extensible)
- Orders
  - Order state (to handle concurrency)
- Inventory
- VendingMachine
  - orchestrator

### Design classes

- Create object and basic structure for entities

### Implementation
- make app working first 
- then think about concurrency handling and other edge cases

### Todos
- app doesn't support multiple notes and coins denominations. need to add that capability
- app doesn't support returning change to user
- add concurrency support