package control

var migration_head = `<?php
use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;
`

//class CreateXnewsTable extends Migration

var migration_head2 = `{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {`

//Schema::create('xnews', function (Blueprint $table) {

var migration_head4 = `
            $table->id();
`

var migration_end = `
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {`

//Schema::dropIfExists('xnews');

var migration_end1 = `
    }
}
`
