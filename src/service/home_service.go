package service

import "fmt"
import "net/http"
import "sudoku"

//import "html/template"

type HomeService struct {
}

func NewHomeService() *HomeService {
	return &HomeService{}
}

var tmpl = `<html>
<head>
<title>Sudoku Puzzle Game</title>

<style>
	.cell {
		width: 25px;
		text-align: center;
	}
</style>

<script src="http://cdn.bootcss.com/jquery/3.1.1/jquery.min.js"></script>

<script>
	var base_url = '%s';

	var set_sudoku = function(puz, init) {
		if (!puz || puz.length != 81) {
			alert("invalid puzzle string");
			return;
		}
		for (var i = 0; i < 81; i++) {
			if (puz[i] != '0') {
				$('#c'+i)[0].value = puz[i];
				if (init) {
					$('#c'+i)[0].disabled = true;
				}
			}
		}
	}

	var random_puzzle = function() {
		var url = base_url + '/puzzles?rand=1';
		console.log("url: " + url);
		$.get(url, function(data) {
			console.log(data);
			var resp = $.parseJSON(data);
			if (resp.status == 'found') {
				var puz = resp.result[0].puz;
				console.log(puz);
				set_sudoku(puz, true);
			}
		});
	}
	
	var get_sudoku = function() {
		var puz = "";
		for (var i = 0; i < 81; i++) {
			puz += $('#c'+i)[0].value || '0';
		}
		return puz;
	};

	var check_puzzle = function() {
		var puz = get_sudoku();
		var url = base_url + '/checker?puz=' + puz;
		$.get(url, function(data) {
			console.log(data);
			var resp = $.parseJSON(data);
			$('#result')[0].innerText = data;
			if (resp.valid == true && resp.blanks == 0) {
				console.log("success");
				alert("Great! you solved this puzzle!");
			}
		});
	};

	var on_page_load = function() {
		console.log("page loaded");
		random_puzzle();
	};
</script>
</head>
<body>
	<div id="main-area" align="center">
		<div id="puzzle" align="center">
		%s
		</div>
		<input type="button" value="Check" onclick="check_puzzle()"> <br/>
		<div id="result"><div> <br/>
	</div>

	<script type="text/javascript">
		on_page_load();
	</script>
</body>
</html>
`

func (hs *HomeService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	table := "<table>\n"
	for i := 0; i < sudoku.Size; i++ {
		table += "\t<tr>\n"
		for j := 0; j < sudoku.Size; j++ {
			table += fmt.Sprintf("\t\t<td><input class=\"cell\" id=\"c%d\"/></td>\n", i*sudoku.Size+j)
		}
		table += "\t</tr>\n"
	}
	table += "</table>\n"

	fmt.Fprintf(writer, tmpl, "http://"+request.Host, table)
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
}
