using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace StudyBud.Data.Migrations
{
    /// <inheritdoc />
    public partial class AllModelAddv63 : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<byte[]>(
                name: "Content",
                table: "Syllabi",
                type: "BLOB",
                nullable: false,
                defaultValue: new byte[0]);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "Content",
                table: "Syllabi");
        }
    }
}
