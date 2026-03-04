# Migration Guide: v6.0.0

## Breaking Changes

### Metadata values must be strings

**Why**: The GoCardless API only accepts string values for metadata, but the Go types previously defined the type as `map[string]interface{}`. This caused confusing runtime errors. v6.0.0 fixes the types to match the API.

**Impact**: Code that passes non-string values to `Metadata` fields will fail to compile.

---

## Quick Migration

### Option 1: Use Helper Functions

The easiest way to migrate is using the new `ToMetadata()` helper:

```go
import "github.com/gocardless/gocardless-pro-go/v6"

// ❌ BEFORE (v6.x) - Compiled but failed at runtime
client.Customers.Create(ctx, gocardless.CustomerCreateParams{
    Email: "user@example.com",
    Metadata: map[string]interface{}{
        "user_id": 12345,           // int
        "is_active": true,          // bool
        "tags": []string{"vip"},    // slice
    },
})

// ✅ AFTER (v6.0.0) - One function call
client.Customers.Create(ctx, gocardless.CustomerCreateParams{
    Email: "user@example.com",
    Metadata: gocardless.ToMetadata(map[string]interface{}{
        "user_id": 12345,           // Auto-converts to "12345"
        "is_active": true,          // Auto-converts to "true"
        "tags": []string{"vip"},    // Auto-converts to `["vip"]`
    }),
})
```

### Option 2: Manual Conversion

If you prefer explicit control:

```go
client.Customers.Create(ctx, gocardless.CustomerCreateParams{
    Email: "user@example.com",
    Metadata: map[string]string{
        "user_id": strconv.Itoa(12345),              // "12345"
        "is_active": strconv.FormatBool(true),       // "true"
        "tags": `["vip"]`,                           // JSON string
    },
})
```

---

## Helper Functions Reference

### `ToMetadata(obj)`

Converts a map with mixed value types to metadata format:

```go
metadata := gocardless.ToMetadata(map[string]interface{}{
    "user_id": 12345,
    "is_premium": true,
    "signup_date": time.Now(),
    "preferences": map[string]string{"theme": "dark"},
})

// Result: map[string]string{
//   "user_id": "12345",
//   "is_premium": "true",
//   "signup_date": "2024-01-15 10:30:00...",
//   "preferences": `{"theme":"dark"}`,
// }
```

### `ToMetadataValue(value)`

Converts a single value:

```go
gocardless.ToMetadataValue(12345)                    // "12345"
gocardless.ToMetadataValue(true)                     // "true"
gocardless.ToMetadataValue([]string{"vip", "pro"})  // `["vip","pro"]`
gocardless.ToMetadataValue(map[string]string{...})  // JSON string
```

### `IsValidMetadata(obj)`

Check if metadata is valid:

```go
if gocardless.IsValidMetadata(metadata) {
    // All values are strings
    client.Customers.Create(ctx, gocardless.CustomerCreateParams{
        Metadata: metadata.(map[string]string),
    })
}
```

### `ParseMetadataValue(value, target)`

Parse metadata values back to their original types:

```go
customer, _ := client.Customers.Get(ctx, "CU123")

// Parse back to original types
var userID int
gocardless.ParseMetadataValue(customer.Metadata["user_id"], &userID)  // 12345

var isActive bool
gocardless.ParseMetadataValue(customer.Metadata["is_active"], &isActive)  // true

var tags []string
gocardless.ParseMetadataValue(customer.Metadata["tags"], &tags)  // []string{"vip"}
```

---

## Common Migration Patterns

### Pattern 1: Numeric IDs

```go
// ❌ Before
Metadata: map[string]interface{}{
    "user_id": userID,
}

// ✅ After (manual)
Metadata: map[string]string{
    "user_id": strconv.FormatInt(userID, 10),
}

// ✅ After (helper)
Metadata: gocardless.ToMetadata(map[string]interface{}{
    "user_id": userID,
})
```

### Pattern 2: Boolean Flags

```go
// ❌ Before
Metadata: map[string]interface{}{
    "is_premium": user.IsPremium,
}

// ✅ After (manual)
Metadata: map[string]string{
    "is_premium": strconv.FormatBool(user.IsPremium),
}

// ✅ After (helper)
Metadata: gocardless.ToMetadata(map[string]interface{}{
    "is_premium": user.IsPremium,
})
```

### Pattern 3: Slices/Arrays

```go
// ❌ Before
Metadata: map[string]interface{}{
    "tags": []string{"vip", "early_adopter"},
}

// ✅ After (manual)
import "encoding/json"
tagsJSON, _ := json.Marshal([]string{"vip", "early_adopter"})
Metadata: map[string]string{
    "tags": string(tagsJSON),
}

// ✅ After (helper)
Metadata: gocardless.ToMetadata(map[string]interface{}{
    "tags": []string{"vip", "early_adopter"},
})
```

### Pattern 4: Nested Structs/Maps

```go
type Preferences struct {
    Theme string `json:"theme"`
    Lang  string `json:"lang"`
}

// ❌ Before
Metadata: map[string]interface{}{
    "preferences": Preferences{Theme: "dark", Lang: "en"},
}

// ✅ After (manual)
prefsJSON, _ := json.Marshal(Preferences{Theme: "dark", Lang: "en"})
Metadata: map[string]string{
    "preferences": string(prefsJSON),
}

// ✅ After (helper)
Metadata: gocardless.ToMetadata(map[string]interface{}{
    "preferences": Preferences{Theme: "dark", Lang: "en"},
})
```

### Pattern 5: Conditional Values

```go
// ❌ Before
Metadata: map[string]interface{}{
    "referral_code": getReferralCode(), // might return nil
}

// ✅ After (manual)
metadata := make(map[string]string)
if code := getReferralCode(); code != nil {
    metadata["referral_code"] = fmt.Sprintf("%v", code)
}

// ✅ After (helper)
Metadata: gocardless.ToMetadata(map[string]interface{}{
    "referral_code": getReferralCode(), // ToMetadata handles nil
})
```
