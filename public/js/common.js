
function ticTacToeConfirm( message, callback ) {
  if ( typeof message == 'undefined' ) {
    message = 'Confirm';
  }
  if ( typeof callback == 'undefined' ) {
    callback = function(b) {};
  }
  var b = false;

  var defer = new jQuery.Deferred();

  $( '#dialog-confirm p span.message' ).text( message );
  var d = $( '#dialog-confirm' ).dialog({
     option : true,
     resizable: false,
     height:140,
     modal: true,
     buttons: {
       "Ok": function() {
         b = true;
         $( this ).dialog( 'close' );
         defer.resolve();
       },
       "Cancel" : function() {
         b = false;
         $( this ).dialog( 'close' );
         defer.reject();
       }
     }
  });

  d.on( "dialogclose", function( ev, ui ) {
     d.off( "dialogclose" );
     callback(b);
  } );
  $( '#dialog-confirm' ).dialog( 'open' );
  return defer.promise();
}

function ticTacToeAlert( message, callback ) {
  if ( typeof message == "undefined" ) {
    message = "Alert";
  }
  if ( typeof callback == "undefined" ) {
    callback = function() {};
  }

  var returning = false;
  var defer = new jQuery.Deferred();
  $( '#dialog-alert p.message' ).text( message );
  var d = $( '#dialog-alert' ).dialog({
    modal: true,
    buttons: {
      "Ok": function() {
        d.dialog( "close" );
      }
    }
  });
  d.on( "dialogclose", function( ev, ui ) {
      d.off( "dialogclose" );
      callback();
      defer.resolve();
  } );
  d.dialog( 'open' );
  return defer.promise();
}
