package tempconv

// CToF converte uma temperatura em Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converte uma temperatura em Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converte uma temperatura em Kelvin para Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converte uma temperatura em Kelvin para Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
