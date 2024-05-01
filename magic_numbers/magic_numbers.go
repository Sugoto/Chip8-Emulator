package magic_numbers

const (
    PCStart = 0x200
    OpcodeMask = 0xF000
    SubMask = 0x0FFF
    RegXMask = 0x0F00
    RegYMask = 0x00F0
    LastNibbleMask = 0x000F
    LastByteMask = 0x00FF
    FontSetStart = 0x50
    DisplayWidth = 64
    DisplayHeight = 32
    MemorySize = 4096
    NumRegisters = 16
    StackSize = 16
)