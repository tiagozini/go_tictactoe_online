  var userId = 0;

  function setupJqueryDialogs() {
    $( '#dialog-confirm, #dialog-alert' ).dialog({
      autoOpen: false
    });
  }

  var bDecideRematch = true;

  function showAlertMessagesAndKickOut() {
    ticTacToeAlert(
      $( '#message > div' ).text(),
      function() {
        window.location = '/App/Index';
      }
    );
  }

  function bringUserTo( url, x, y ) {
    var extra = "";
    extra += ( typeof x != "undefined" ?  "&x=" + x : "" );
    extra += ( typeof y != "undefined" ?  "&y=" + y : "" );
    window.location = url + '?userId=' + userId + extra;
  }

  function applyJQueryEvents() {
    $( '.playing .box' ).on( 'click' , function() {
      if ( $(this).html() === '' ) {
        var arr = $(this).data( 'coordenates' ).split( ':' );
        bringUserTo('/Refresh/Play', arr[0], arr[1]);
      } else {
        ticTacToeAlert( 'Place already used!' );
      }
    });

    $( '.waiting .box, .finished .box' ).on( 'click' , function() {
      $.when( ticTacToeAlert( $( '#message > div' ).text(), function() { }, false ) );
    });

    if ( $( '#challengerNamesValue' ).val() != '' ) {
      challengerNames = $( '#challengerNamesValue' ).val().split('|');
      $("#challengerNamesText").html("('" + challengerNames[0] + "' VS '" +
        challengerNames[1] + "')");
    }
    setupJqueryDialogs();
  }

  function decideRedirect() {
    if ( $('table').hasClass() == 'waiting_rematch' ) {
      agreeWithRematch();
    }
    if ( $('input#refresh').val() == '1' ) {
      bringUserTo('/Refresh/Room')
    }
    if ( $( 'table' ).hasClass( 'waiting_destruction' ) ) {
      showAlertMessagesAndKickOut();
      return true;
    }
    return false;
  }

  function agreeWithRematch() {
    bringUserTo('/Refresh/AskRematch');
  }

  function decideRematch() {
    bDecideRematch = false;
    ticTacToeConfirm(
      'Your mate ask for rematch. You agree?',
      function(decision) {
        if ( decision ) {
          agreeWithRematch();
        } else {
          bringUserTo('/Refresh/RejectRematch')
        }
        bDecideRematch = true;
      }
    );
  }

  function runOnStart() {
    if ( $( 'table' ).hasClass( 'waiting_rematch_answer' ) ) {
      if ( bDecideRematch ) {
        decideRematch();
      }
    }
  }

  function init() {
    applyJQueryEvents()
    runOnStart();
  }
