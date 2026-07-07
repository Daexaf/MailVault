# Package di Go

Package adalah wadah yang berisi:

- Struct
- Interface
- Function
- Variable
- Constant

Contoh:

package logger

func New()

Package logger tidak memiliki struct Logger.

Function New() mengembalikan \*slog.Logger.

Karena itu Bootstrap menyimpan:

\*slog.Logger

bukan

\*logger.Logger
