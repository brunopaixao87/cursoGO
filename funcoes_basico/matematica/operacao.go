package matematica

/*
  Somar retorna a soma de dois valores inteiros
 */
func Somar(valorEntrada1 int, valorEntrada2 int) int  {
	return valorEntrada1 + valorEntrada2
}

/*
  Multiplicar retorna a multiplicação de dois valores inteiros
 */
func Multiplicar(valorEntrada1 int, valorEntrada2 int) int  {
	return valorEntrada1 * valorEntrada2
}

/*
 Dividir retorna a divisão de dois valores inteiros
 */
func Dividir(valorEntrada1 int, valorEntrada2 int) (resultado int){
	resultado = valorEntrada1 / valorEntrada2
	return
}

/*
 Dividir retorna a divisão e o resto de dois valores inteiros
 */
func DividirComResto(valorEntrada1 int, valorEntrada2 int) (resultado int, resto int){
	resultado = valorEntrada1 / valorEntrada2
	resto = valorEntrada1 % valorEntrada2
	return
}
