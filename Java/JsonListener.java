// Generated from Json.g4 by ANTLR 4.8
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link JsonParser}.
 */
public interface JsonListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link JsonParser#bool}.
	 * @param ctx the parse tree
	 */
	void enterBool(JsonParser.BoolContext ctx);
	/**
	 * Exit a parse tree produced by {@link JsonParser#bool}.
	 * @param ctx the parse tree
	 */
	void exitBool(JsonParser.BoolContext ctx);
	/**
	 * Enter a parse tree produced by {@link JsonParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(JsonParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link JsonParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(JsonParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link JsonParser#pair}.
	 * @param ctx the parse tree
	 */
	void enterPair(JsonParser.PairContext ctx);
	/**
	 * Exit a parse tree produced by {@link JsonParser#pair}.
	 * @param ctx the parse tree
	 */
	void exitPair(JsonParser.PairContext ctx);
	/**
	 * Enter a parse tree produced by {@link JsonParser#object}.
	 * @param ctx the parse tree
	 */
	void enterObject(JsonParser.ObjectContext ctx);
	/**
	 * Exit a parse tree produced by {@link JsonParser#object}.
	 * @param ctx the parse tree
	 */
	void exitObject(JsonParser.ObjectContext ctx);
	/**
	 * Enter a parse tree produced by {@link JsonParser#array}.
	 * @param ctx the parse tree
	 */
	void enterArray(JsonParser.ArrayContext ctx);
	/**
	 * Exit a parse tree produced by {@link JsonParser#array}.
	 * @param ctx the parse tree
	 */
	void exitArray(JsonParser.ArrayContext ctx);
	/**
	 * Enter a parse tree produced by {@link JsonParser#init}.
	 * @param ctx the parse tree
	 */
	void enterInit(JsonParser.InitContext ctx);
	/**
	 * Exit a parse tree produced by {@link JsonParser#init}.
	 * @param ctx the parse tree
	 */
	void exitInit(JsonParser.InitContext ctx);
}