package grafico

import (
	"fmt"
	"os"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"experimentoMatriz/benchmark"
)

func GerarGraficos(resultado benchmark.ResultadosBenchmark){
	var ordensStr []string
	var dataBaseline,dataReforged []opts.BarData
	var barra1,barra2,barra3,barra4 *charts.Bar
	var diferencaPercent,tempoBase,tempoRef float64
	var diferencaPercentBar []opts.BarData
	var arquivo *os.File
	var cont,ordem,tempo int 
	var erro error
	
	//debug
	fmt.Println("\n--- DADOS RECEBIDOS PELA FUNÇÃO DE GRÁFICOS ---")
	fmt.Printf("Tempos Reforged recebidos: %v\n", resultado.TempoReforged)
	fmt.Println("-------------------------------------------------")


	ordensStr=make([]string,len(resultado.Ordens))
	for cont,ordem=range resultado.Ordens{
		ordensStr[cont]=fmt.Sprintf("%d",ordem)
	}
	
	//Grafico 1-Baseline
	barra1=charts.NewBar()
	barra1.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Tempo Médio (baseline)"}),
		charts.WithYAxisOpts(opts.YAxis{Type: "log"}),
	)
	//converter os tempos baseline em dados de barro(chamado BarData)
	dataBaseline=make([]opts.BarData,len(resultado.TempoBaseline))
	for cont,tempo=range resultado.TempoBaseline{
		dataBaseline[cont]=opts.BarData{Value:(tempo)}
	}
	barra1.SetXAxis(ordensStr).AddSeries("Baseline",dataBaseline)

	//Gráfico 2-otimizado
	barra2=charts.NewBar()
	barra2.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Tempo Médio (reforged)"}),
		charts.WithYAxisOpts(opts.YAxis{Type: "log"}),
	)
	//converter o tempo em bardata 
	dataReforged=make([]opts.BarData,len(resultado.TempoReforged))
	for cont,tempo=range resultado.TempoReforged {
		dataReforged[cont]=opts.BarData{Value: (tempo)}
	}
	barra2.SetXAxis(ordensStr).AddSeries("Otimizado",dataReforged)

	//comparação lado a lado 
	barra3=charts.NewBar()
	barra3.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Comparação:baseline X reforged"}),
		charts.WithYAxisOpts(opts.YAxis{Type: "log"}),
	)
	barra3.SetXAxis(ordensStr).
		AddSeries("Baseline",dataBaseline).
		AddSeries("Otimizado",dataReforged)
	
	//diferença percentual 
	barra4=charts.NewBar()
	barra4.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Diferença Percentual (BASELINE X REFORGED)"}),
	)

	//var diferencaPercentBar []opts.BarData 
		for cont=range resultado.TempoBaseline{
		tempoBase=float64(resultado.TempoBaseline[cont])
		tempoRef=float64(resultado.TempoReforged[cont])

		diferencaPercent=(tempoBase-tempoRef)/tempoBase*100
		diferencaPercentBar=append(diferencaPercentBar,opts.BarData{Value: diferencaPercent})
	}
	
	barra4.SetXAxis(ordensStr).AddSeries("Diferença %",diferencaPercentBar)
	arquivo,erro=os.Create("graficosExperimento.html")
	if erro!=nil{
		fmt.Println("Erro ao criar arquivo HTML: ",erro)
	}

	//renderizar graficos 
	barra1.Render(arquivo)
	barra2.Render(arquivo)
	barra3.Render(arquivo)
	barra4.Render(arquivo)
}
