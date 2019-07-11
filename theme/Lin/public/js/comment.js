$(document).ready(function() {
    /**
     * 回复框
     */
    $('.reply-form')
        .transition('hide')
    ;
    $('.reply').click(function(){
        var id = $(this).attr("id");
        $('#' + id + '-reply')
            .transition('slide down')
        ;
    });

    /**
     * TODO 临时遮罩层
     */
    $('.linshi-dimmer')
        .dimmer({
            on: 'hover'
        })
    ;
});
