using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace StudyBud.Data.Migrations
{
    /// <inheritdoc />
    public partial class AllModelAddv61 : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Cohorts_Users_UserId",
                table: "Cohorts");

            migrationBuilder.DropForeignKey(
                name: "FK_Degrees_Users_UserId",
                table: "Degrees");

            migrationBuilder.DropForeignKey(
                name: "FK_Schools_Users_UserId",
                table: "Schools");

            migrationBuilder.DropForeignKey(
                name: "FK_Syllabi_Users_UserId",
                table: "Syllabi");

            migrationBuilder.DropTable(
                name: "Users");

            migrationBuilder.AddColumn<string>(
                name: "Address",
                table: "AspNetUsers",
                type: "TEXT",
                nullable: true);

            migrationBuilder.AddColumn<string>(
                name: "FName",
                table: "AspNetUsers",
                type: "TEXT",
                nullable: true);

            migrationBuilder.AddColumn<float>(
                name: "GPA",
                table: "AspNetUsers",
                type: "REAL",
                nullable: true);

            migrationBuilder.AddColumn<string>(
                name: "LName",
                table: "AspNetUsers",
                type: "TEXT",
                nullable: true);

            migrationBuilder.AddColumn<string>(
                name: "MInitial",
                table: "AspNetUsers",
                type: "TEXT",
                nullable: true);

            migrationBuilder.AddColumn<string>(
                name: "Name",
                table: "AspNetUsers",
                type: "TEXT",
                nullable: true);

            migrationBuilder.AddColumn<bool>(
                name: "Subscribed",
                table: "AspNetUsers",
                type: "INTEGER",
                nullable: true);

            migrationBuilder.AddForeignKey(
                name: "FK_Cohorts_AspNetUsers_UserId",
                table: "Cohorts",
                column: "UserId",
                principalTable: "AspNetUsers",
                principalColumn: "Id");

            migrationBuilder.AddForeignKey(
                name: "FK_Degrees_AspNetUsers_UserId",
                table: "Degrees",
                column: "UserId",
                principalTable: "AspNetUsers",
                principalColumn: "Id");

            migrationBuilder.AddForeignKey(
                name: "FK_Schools_AspNetUsers_UserId",
                table: "Schools",
                column: "UserId",
                principalTable: "AspNetUsers",
                principalColumn: "Id");

            migrationBuilder.AddForeignKey(
                name: "FK_Syllabi_AspNetUsers_UserId",
                table: "Syllabi",
                column: "UserId",
                principalTable: "AspNetUsers",
                principalColumn: "Id");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Cohorts_AspNetUsers_UserId",
                table: "Cohorts");

            migrationBuilder.DropForeignKey(
                name: "FK_Degrees_AspNetUsers_UserId",
                table: "Degrees");

            migrationBuilder.DropForeignKey(
                name: "FK_Schools_AspNetUsers_UserId",
                table: "Schools");

            migrationBuilder.DropForeignKey(
                name: "FK_Syllabi_AspNetUsers_UserId",
                table: "Syllabi");

            migrationBuilder.DropColumn(
                name: "Address",
                table: "AspNetUsers");

            migrationBuilder.DropColumn(
                name: "FName",
                table: "AspNetUsers");

            migrationBuilder.DropColumn(
                name: "GPA",
                table: "AspNetUsers");

            migrationBuilder.DropColumn(
                name: "LName",
                table: "AspNetUsers");

            migrationBuilder.DropColumn(
                name: "MInitial",
                table: "AspNetUsers");

            migrationBuilder.DropColumn(
                name: "Name",
                table: "AspNetUsers");

            migrationBuilder.DropColumn(
                name: "Subscribed",
                table: "AspNetUsers");

            migrationBuilder.CreateTable(
                name: "Users",
                columns: table => new
                {
                    UserId = table.Column<string>(type: "TEXT", nullable: false),
                    Address = table.Column<string>(type: "TEXT", nullable: true),
                    Email = table.Column<string>(type: "TEXT", nullable: true),
                    FName = table.Column<string>(type: "TEXT", nullable: true),
                    GPA = table.Column<float>(type: "REAL", nullable: true),
                    LName = table.Column<string>(type: "TEXT", nullable: true),
                    MInitial = table.Column<string>(type: "TEXT", nullable: true),
                    Name = table.Column<string>(type: "TEXT", nullable: true),
                    Phone = table.Column<string>(type: "TEXT", nullable: true),
                    Subscribed = table.Column<bool>(type: "INTEGER", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Users", x => x.UserId);
                });

            migrationBuilder.AddForeignKey(
                name: "FK_Cohorts_Users_UserId",
                table: "Cohorts",
                column: "UserId",
                principalTable: "Users",
                principalColumn: "UserId");

            migrationBuilder.AddForeignKey(
                name: "FK_Degrees_Users_UserId",
                table: "Degrees",
                column: "UserId",
                principalTable: "Users",
                principalColumn: "UserId");

            migrationBuilder.AddForeignKey(
                name: "FK_Schools_Users_UserId",
                table: "Schools",
                column: "UserId",
                principalTable: "Users",
                principalColumn: "UserId");

            migrationBuilder.AddForeignKey(
                name: "FK_Syllabi_Users_UserId",
                table: "Syllabi",
                column: "UserId",
                principalTable: "Users",
                principalColumn: "UserId");
        }
    }
}
