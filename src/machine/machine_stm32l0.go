// +build stm32l0

package machine

// Peripheral abstraction layer for the stm32l0

import (
	"device/stm32"
	"runtime/interrupt"
	"unsafe"
)

func CPUFrequency() uint32 {
	return 32000000
}

const (
	PA0  = portA + 0
	PA1  = portA + 1
	PA2  = portA + 2
	PA3  = portA + 3
	PA4  = portA + 4
	PA5  = portA + 5
	PA6  = portA + 6
	PA7  = portA + 7
	PA8  = portA + 8
	PA9  = portA + 9
	PA10 = portA + 10
	PA11 = portA + 11
	PA12 = portA + 12
	PA13 = portA + 13
	PA14 = portA + 14
	PA15 = portA + 15

	PB0  = portB + 0
	PB1  = portB + 1
	PB2  = portB + 2
	PB3  = portB + 3
	PB4  = portB + 4
	PB5  = portB + 5
	PB6  = portB + 6
	PB7  = portB + 7
	PB8  = portB + 8
	PB9  = portB + 9
	PB10 = portB + 10
	PB11 = portB + 11
	PB12 = portB + 12
	PB13 = portB + 13
	PB14 = portB + 14
	PB15 = portB + 15

	PC0  = portC + 0
	PC1  = portC + 1
	PC2  = portC + 2
	PC3  = portC + 3
	PC4  = portC + 4
	PC5  = portC + 5
	PC6  = portC + 6
	PC7  = portC + 7
	PC8  = portC + 8
	PC9  = portC + 9
	PC10 = portC + 10
	PC11 = portC + 11
	PC12 = portC + 12
	PC13 = portC + 13
	PC14 = portC + 14
	PC15 = portC + 15

	PD0  = portD + 0
	PD1  = portD + 1
	PD2  = portD + 2
	PD3  = portD + 3
	PD4  = portD + 4
	PD5  = portD + 5
	PD6  = portD + 6
	PD7  = portD + 7
	PD8  = portD + 8
	PD9  = portD + 9
	PD10 = portD + 10
	PD11 = portD + 11
	PD12 = portD + 12
	PD13 = portD + 13
	PD14 = portD + 14
	PD15 = portD + 15

	PE0  = portE + 0
	PE1  = portE + 1
	PE2  = portE + 2
	PE3  = portE + 3
	PE4  = portE + 4
	PE5  = portE + 5
	PE6  = portE + 6
	PE7  = portE + 7
	PE8  = portE + 8
	PE9  = portE + 9
	PE10 = portE + 10
	PE11 = portE + 11
	PE12 = portE + 12
	PE13 = portE + 13
	PE14 = portE + 14
	PE15 = portE + 15

	PH0 = portH + 0
	PH1 = portH + 1
)

func (p Pin) getPort() *stm32.GPIO_Type {
	switch p / 16 {
	case 0:
		return stm32.GPIOA
	case 1:
		return stm32.GPIOB
	case 2:
		return stm32.GPIOC
	case 3:
		return stm32.GPIOD
	case 4:
		return stm32.GPIOE
	case 7:
		return stm32.GPIOH
	default:
		panic("machine: unknown port")
	}
}

// enableClock enables the clock for this desired GPIO port.
func (p Pin) enableClock() {
	switch p / 16 {
	case 0:
		stm32.RCC.IOPENR.SetBits(stm32.RCC_IOPENR_IOPAEN)
	case 1:
		stm32.RCC.IOPENR.SetBits(stm32.RCC_IOPENR_IOPBEN)
	case 2:
		stm32.RCC.IOPENR.SetBits(stm32.RCC_IOPENR_IOPCEN)
	case 3:
		stm32.RCC.IOPENR.SetBits(stm32.RCC_IOPENR_IOPDEN)
	case 4:
		stm32.RCC.IOPENR.SetBits(stm32.RCC_IOPENR_IOPEEN)
	case 7:
		stm32.RCC.IOPENR.SetBits(stm32.RCC_IOPENR_IOPHEN)
	default:
		panic("machine: unknown port")
	}
}

// Enable peripheral clock
func enableAltFuncClock(bus unsafe.Pointer) {
	switch bus {
	case unsafe.Pointer(stm32.DAC): // DAC interface clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_DACEN)
	case unsafe.Pointer(stm32.PWR): // Power interface clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_PWREN)
	case unsafe.Pointer(stm32.I2C3): // I2C3 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_I2C3EN)
	case unsafe.Pointer(stm32.I2C2): // I2C2 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_I2C2EN)
	case unsafe.Pointer(stm32.I2C1): // I2C1 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_I2C1EN)
	case unsafe.Pointer(stm32.USART5): // UART5 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_USART5EN)
	case unsafe.Pointer(stm32.USART4): // UART4 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_USART4EN)
	case unsafe.Pointer(stm32.USART2): // USART2 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_USART2EN)
	case unsafe.Pointer(stm32.SPI2): // SPI2 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_SPI2EN)
	case unsafe.Pointer(stm32.LPUART1): // LPUART1 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_LPUART1EN)
	case unsafe.Pointer(stm32.WWDG): // Window watchdog clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_WWDGEN)
	case unsafe.Pointer(stm32.TIM7): // TIM7 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_TIM7EN)
	case unsafe.Pointer(stm32.TIM6): // TIM6 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_TIM6EN)
	case unsafe.Pointer(stm32.TIM3): // TIM3 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_TIM3EN)
	case unsafe.Pointer(stm32.TIM2): // TIM2 clock enable
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_TIM2EN)
	case unsafe.Pointer(stm32.SYSCFG): // System configuration controller clock enable
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_SYSCFGEN)
	case unsafe.Pointer(stm32.SPI1): // SPI1 clock enable
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_SPI1EN)
	case unsafe.Pointer(stm32.ADC): // ADC clock enable
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_ADCEN)
	case unsafe.Pointer(stm32.USART1): // USART1 clock enable
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_USART1EN)
	}
}

//---------- UART related types and code

// UART representation
type UART struct {
	Buffer          *RingBuffer
	Bus             *stm32.USART_Type
	Interrupt       interrupt.Interrupt
	AltFuncSelector uint8
}

// Configure the UART.
func (uart UART) configurePins(config UARTConfig) {
	// enable the alternate functions on the TX and RX pins
	config.TX.ConfigureAltFunc(PinConfig{Mode: PinModeUARTTX}, uart.AltFuncSelector)
	config.RX.ConfigureAltFunc(PinConfig{Mode: PinModeUARTRX}, uart.AltFuncSelector)
}

// UART baudrate calc based on the bus and clockspeed
func (uart UART) getBaudRateDivisor(baudRate uint32) uint32 {
	var clock, rate uint32
	switch uart.Bus {
	case stm32.LPUART1:
		clock = CPUFrequency() / 2 // APB1 Frequency
		rate = uint32((256 * clock) / baudRate)
	case stm32.USART1:
		clock = CPUFrequency() / 2 // APB2 Frequency
		rate = uint32(clock / baudRate)
	case stm32.USART2:
		clock = CPUFrequency() / 2 // APB1 Frequency
		rate = uint32(clock / baudRate)
	}

	return rate
}

//---------- SPI related types and code

// SPI on the STM32Fxxx using MODER / alternate function pins
type SPI struct {
	Bus             *stm32.SPI_Type
	AltFuncSelector uint8
}

// Set baud rate for SPI
func (spi SPI) getBaudRate(config SPIConfig) uint32 {
	var conf uint32

	localFrequency := config.Frequency

	// Default
	if config.Frequency == 0 {
		config.Frequency = 4e6
	}

	if spi.Bus != stm32.SPI1 {
		// Assume it's SPI2 or SPI3 on APB1 at 1/2 the clock frequency of APB2, so
		//  we want to pretend to request 2x the baudrate asked for
		localFrequency = localFrequency * 2
	}

	// set frequency dependent on PCLK prescaler. Since these are rather weird
	// speeds due to the CPU freqency, pick a range up to that frquency for
	// clients to use more human-understandable numbers, e.g. nearest 100KHz

	// These are based on APB2 clock frquency (84MHz on the discovery board)
	// TODO: also include the MCU/APB clock setting in the equation
	switch {
	case localFrequency < 328125:
		conf = stm32.SPI_CR1_BR_Div256
	case localFrequency < 656250:
		conf = stm32.SPI_CR1_BR_Div128
	case localFrequency < 1312500:
		conf = stm32.SPI_CR1_BR_Div64
	case localFrequency < 2625000:
		conf = stm32.SPI_CR1_BR_Div32
	case localFrequency < 5250000:
		conf = stm32.SPI_CR1_BR_Div16
	case localFrequency < 10500000:
		conf = stm32.SPI_CR1_BR_Div8
		// NOTE: many SPI components won't operate reliably (or at all) above 10MHz
		// Check the datasheet of the part
	case localFrequency < 21000000:
		conf = stm32.SPI_CR1_BR_Div4
	case localFrequency < 42000000:
		conf = stm32.SPI_CR1_BR_Div2
	default:
		// None of the specific baudrates were selected; choose the lowest speed
		conf = stm32.SPI_CR1_BR_Div256
	}

	return conf << stm32.SPI_CR1_BR_Pos
}

// Configure SPI pins for input output and clock
func (spi SPI) configurePins(config SPIConfig) {
	config.SCK.ConfigureAltFunc(PinConfig{Mode: PinModeSPICLK}, spi.AltFuncSelector)
	config.SDO.ConfigureAltFunc(PinConfig{Mode: PinModeSPISDO}, spi.AltFuncSelector)
	config.SDI.ConfigureAltFunc(PinConfig{Mode: PinModeSPISDI}, spi.AltFuncSelector)
}
