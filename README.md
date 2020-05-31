# Lab - Organizador de arquivos em GoLang

Projeto para organizar arquivos em GoLang para estudos.

## Introdução

Essas instruções fornecerão uma cópia do projeto em execução na sua máquina local para fins de desenvolvimento e teste.

### Prerequisitos

O que você precisa para baixar, rodar e disponibilizar.

* SDK Go - versão utilizada 1.14.3
* IDE de sua preferência 

### Instalação

Após a execução do pre requisitos, segue um passo a passo de como rodar localmente.

Clonar o repositório

```
git clone git@github.com:robsonpedroso/lab-go-organizer.git
```

Acesse a pasta do projeto, exemplo:

```
cd lab-go-organizer
```

Execute o comando **go build** para compilar e gerar o arquivo executável

```
go build organizer.go
```

Ou execute o comando **go run** para apenas executar o programa

```
go run organizer.go
```

### Parâmetros

Existem 3 parâmetros no projeto:
* **s** - Serve para indicar a pasta de origem (source). Ex.: c:\fotos\
* **o** - Serve para indicar a pasta de destino (output). Ex.: f:\fotos-organizadas\
* **m** - Serve para indicar se será gerado a saida com o diretório do mês. Por padrão é **true**. Valores permitidos (**true** ou **false**)

Se não for passado nenhum parâmetro, será solicitado o diretório de origem e de destino.

## Publicação

Não foi publicado

## Autores

* **Robson Pedroso** - *Projeto inicial* - [RobsonPedroso](https://github.com/robsonpedroso)

## License

[MIT](https://github.com/robsonpedroso/lab-go-organizer/blob/master/LICENSE)

## Ferramentas

* [Golang](https://golang.org/)
