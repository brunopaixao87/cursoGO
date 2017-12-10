package matematica

/*
 Calculo Dado uma função e de dois inteiros retorna o calculos desse números
 */
func Calcular(calculo func(valor1 int, valor2 int) int, valorEntrada1 int, valorEntrada2 int ) int  {
    return calculo(valorEntrada1, valorEntrada2);
}
