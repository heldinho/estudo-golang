#!/bin/bash

# Referencia
# https://www.devmedia.com.br/introducao-ao-shell-script-no-linux/25778

echo "Seu nome de usuário é:"
whoami
echo "Info de hora atual e tempo que o computador está ligado:"
uptime
echo "O script está executando do diretório:"
pwd

# Declarando e utilizando variáveis
site="www.devmedia.com.br"
meu_numero_favorito=15
_cidade="São Paulo"
echo "Um ótimo site para você aprender a programar e se manter  atualizado é: $site"
echo "Meu número favorito é: $meu_numero_favorito"
echo "Minha cidade natal é: $_cidade"

# Atribuindo saídas de comandos a variáveis
system_info=`df -h` # Também poderia ser system_info=$(df -h)
system_info2=$(df -h)
echo "$system_info"
echo "$system_info2"

echo ""
echo ""

# Capturando a entrada de dados do usuário
echo "Qual o seu nome?"
read name
echo "Olá, $name. Bem vindo(a)."

echo ""
echo ""

# Comandos de seleção ou de tomada de decisão
echo "Digite um número qualquer:"
read num
if [ "$num" -gt 20 ]
	then
		echo "Este número é maior que 20!"
	else
		echo "Este número é menor que 20!"
fi